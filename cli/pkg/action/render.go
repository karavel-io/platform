// Copyright 2021 MIKAMAI s.r.l
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package action

import (
	"fmt"
	"github.com/mikamai/karavel/cli/internal/gitutils"
	"github.com/mikamai/karavel/cli/internal/helmw"
	"github.com/mikamai/karavel/cli/internal/plan"
	"github.com/mikamai/karavel/cli/internal/utils"
	"github.com/mikamai/karavel/cli/internal/utils/predicate"
	"github.com/mikamai/karavel/cli/pkg/config"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type RenderParams struct {
	ConfigPath string
	SkipGit    bool
}

func Render(log logger.Logger, params RenderParams) error {
	cpath := params.ConfigPath
	skipGit := params.SkipGit
	workdir := filepath.Dir(cpath)
	vendorDir := filepath.Join(workdir, "vendor")
	appsDir := filepath.Join(workdir, "applications")
	projsDir := filepath.Join(workdir, "projects")
	argoEnabled := true

	log.Infof("Rendering new Karavel project with config file %s", cpath)

	log.Debug("Reading config file")
	cfg, err := config.ReadFrom(log.Writer(), cpath)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}

	log.Debugf("Using channel '%s'", cfg.Channel)
	log.Debugf("Updating Karavel Charts repository %s", cfg.HelmRepoUrl)
	if err := helmw.SetupHelm(cfg.HelmRepoUrl); err != nil {
		return errors.Wrap(err, "failed to setup Karavel Charts repository")
	}

	log.Debug("Creating render plan from config")
	p, err := plan.NewFromConfig(log, &cfg)
	if err != nil {
		return errors.Wrap(err, "failed to instantiate render plan from config")
	}

	log.Debug("Validating render plan")
	if err := p.Validate(); err != nil {
		return err
	}

	argo := p.GetComponent("argocd")
	if argo == nil {
		argoEnabled = false
		log.Warnf("ArgoCD component is missing. GitOps integrations will be disabled")
	}

	assertDirs := []string{vendorDir}
	if argoEnabled {
		assertDirs = append(assertDirs, appsDir, projsDir)
	}

	for _, dir := range assertDirs {
		log.Debugf("Asserting directory %s", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, "failed to create directory %s", dir)
		}
	}

	var wg sync.WaitGroup
	ch := make(chan utils.Pair)

	var apps []string
	var renderDirs []string
	if argoEnabled {
		renderDirs = []string{"applications", "projects"}
	}
	dirInfos, err := ioutil.ReadDir(vendorDir)
	if err != nil {
		return err
	}

	dirs := make(map[string]struct{}, len(dirInfos))
	for _, i := range dirInfos {
		dirs[i.Name()] = struct{}{}
	}

	repoPath, repoUrl := "", ""
	if !skipGit && argoEnabled {
		res := argo.GetParam("git.repo")
		log.Debug("Finding remote git repository URL to configure ArgoCD applications")
		dir, url, err := gitutils.GetOriginRemote(log, workdir, res.String())
		if err != nil {
			return err
		}

		path, err := filepath.Rel(dir, workdir)
		if err != nil {
			return err
		}

		repoPath, repoUrl = path, url
	}

	// empty line for nice logs
	log.Info()

	for _, c := range p.Components() {
		if c.IsBootstrap() {
			renderDirs = append(renderDirs, filepath.Join("vendor", c.Name()))
		}
		delete(dirs, c.Name())

		wg.Add(1)
		go func(comp *plan.Component) {
			defer wg.Done()

			msg := fmt.Sprintf("failed to render component '%s'", comp.Name())
			outdir := filepath.Join(vendorDir, comp.Name())
			log.Infof("Rendering component %s at %s", comp.DebugLabel(), strings.ReplaceAll(outdir, filepath.Dir(workdir)+"/", ""))
			log.Debugf("Component %s params: %s", comp.DebugLabel(), comp.Params())

			if err := comp.Render(log, outdir); err != nil {
				ch <- utils.NewPair(msg, err)
				return
			}

			if argoEnabled {
				log.Debugf("Rendering application manifest for component %s", comp.DebugLabel())
				appFile := comp.Name() + ".yml"
				apps = append(apps, appFile)
				appfile := filepath.Join(appsDir, appFile)
				// if the application file already exists, we skip it. It has already been created
				// and we don't want to overwrite any changes the user may have made
				_, err = os.Stat(appfile)
				if !os.IsNotExist(err) {
					ch <- utils.NewPair(msg, err)
					return
				}

				argoNs := argo.Namespace()
				vendorPath := path.Join(repoPath, "vendor", comp.Name())
				if err := comp.RenderApplication(argoNs, repoUrl, vendorPath, appfile); err != nil {
					ch <- utils.NewPair(msg, err)
				}
			}
		}(c)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for dir := range dirs {
			log.Debugf("deleting extraneous directory '%s' in vendor", dir)
			if err := os.RemoveAll(filepath.Join(vendorDir, dir)); err != nil {
				ch <- utils.NewPair(fmt.Sprintf("failed to delete extraneous directory '%s' in vendor", dir), err)
			}
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for pair := range ch {
		err := pair.ErrorB()
		if err != nil {
			return errors.Wrap(err, pair.StringA())
		}
	}

	if argoEnabled {
		argoNs := argo.Namespace()
		sort.Strings(apps)
		if err := utils.RenderKustomizeFile(appsDir, apps, predicate.IsStringInSlice(apps)); err != nil {
			return errors.Wrap(err, "failed to render applications kustomization.yml")
		}

		infraProj := "infrastructure.yml"
		if err := ioutil.WriteFile(filepath.Join(projsDir, infraProj), []byte(fmt.Sprintf(argoProject, argoNs)), 0655); err != nil {
			return errors.Wrap(err, "failed to render infrastructure project file")
		}

		projs := []string{infraProj}
		if err := utils.RenderKustomizeFile(projsDir, projs, predicate.IsStringInSlice(projs)); err != nil {
			return errors.Wrap(err, "failed to render projects kustomization.yml")
		}
	}

	if err := utils.RenderKustomizeFile(workdir, renderDirs, predicate.StringOr(predicate.IsStringInSlice(renderDirs), predicate.StringHasPrefix("vendor"))); err != nil {
		return errors.Wrap(err, "failed to render kustomization.yml")
	}

	return nil
}

const argoProject = `
apiVersion: argoproj.io/v1alpha1
kind: AppProject
metadata:
  name: infrastructure
  namespace: %s
spec:
  description: Platform infrastructure components
  sourceRepos:
    - '*'
  destinations:
    - namespace: '*'
      server: 'https://kubernetes.default.svc'
  clusterResourceWhitelist:
    - group: '*'
      kind: '*'

`
