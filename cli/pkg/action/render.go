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
	"github.com/mikamai/karavel/cli/internal/helmw"
	"github.com/mikamai/karavel/cli/internal/plan"
	"github.com/mikamai/karavel/cli/internal/utils"
	"github.com/mikamai/karavel/cli/internal/utils/predicate"
	"github.com/mikamai/karavel/cli/pkg/config"
	"github.com/mikamai/karavel/cli/pkg/logger"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type RenderParams struct {
	ConfigPath string
}

func Render(log logger.Logger, params RenderParams) error {
	cpath := params.ConfigPath
	workdir := filepath.Dir(cpath)
	vendorDir := filepath.Join(workdir, "vendor")
	appsDir := filepath.Join(workdir, "applications")
	projsDir := filepath.Join(workdir, "projects")

	log.Infof("Rendering new Karavel project with config file %s", cpath)

	log.Debug("Reading config file")
	cfg, err := config.ReadFrom(log.Writer(), cpath)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}

	log.Debugf("Setting up Karavel Charts repository %s", cfg.HelmRepoUrl)
	if err := helmw.SetupHelm(cfg.HelmRepoUrl); err != nil {
		return errors.Wrap(err, "failed to setup Karavel Charts repository")
	}

	log.Info()

	p, err := plan.NewFromConfig(&cfg)
	if err != nil {
		return errors.Wrap(err, "failed to instantiate render plan from config")
	}

	if err := p.Validate(); err != nil {
		return err
	}

	for _, dir := range []string{appsDir, projsDir} {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return errors.Wrapf(err, "failed to create directory %s", dir)
		}
	}

	argo := p.GetComponent("argocd")
	if argo == nil {
		return errors.New("component 'argocd' is required")
	}
	argoNs := argo.Namespace()

	var wg sync.WaitGroup
	ch := make(chan utils.Pair)
	var apps []string
	renderDirs := []string{"applications", "projects"}
	dirInfos, err := ioutil.ReadDir(vendorDir)
	if err != nil {
		return err
	}

	dirs := make(map[string]struct{}, len(dirInfos))
	for _, i := range dirInfos {
		dirs[i.Name()] = struct{}{}
	}

	for _, c := range p.Components() {
		appFile := c.Name() + ".yml"
		apps = append(apps, appFile)
		if c.IsBootstrap() {
			renderDirs = append(renderDirs, filepath.Join("vendor", c.Name()))
		}
		delete(dirs, c.Name())

		wg.Add(1)
		go func(comp *plan.Component) {
			defer wg.Done()

			msg := fmt.Sprintf("failed to render component '%s'", comp.Name())
			outdir := filepath.Join(vendorDir, comp.Name())
			var alias string
			if name := comp.NameOverride(); name != "" {
				alias = fmt.Sprintf(" with alias '%s'", name)
			}
			log.Infof("Rendering component '%s' %s%s at %s", comp.ComponentName(), comp.Version(), alias, strings.ReplaceAll(outdir, filepath.Dir(workdir)+"/", ""))
			log.Debugf("Component '%s' %s params: %s", comp.Name(), comp.Version(), comp.Params())

			if err := comp.Render(outdir); err != nil {
				ch <- utils.NewPair(msg, err)
				return
			}

			log.Debugf("Rendering application manifest for component '%s' %s", comp.Name(), comp.Version())
			appfile := filepath.Join(appsDir, appFile)
			// if the application file already exists, we skip it. It has already been created
			// and we don't want to overwrite any changes the user may have made
			_, err := os.Stat(appfile)
			if !os.IsNotExist(err) {
				ch <- utils.NewPair(msg, err)
				return
			}

			// TODO: git integration to detect repo and path if not provided in config
			if err := comp.RenderApplication(argoNs, "TODO", "TODO", appfile); err != nil {
				ch <- utils.NewPair(msg, err)
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

	if err := utils.RenderKustomizeFile(workdir, renderDirs, predicate.StringOr(predicate.IsStringInSlice(renderDirs), predicate.StringHasPrefix("vendor"))); err != nil {
		return errors.Wrap(err, "failed to render render kustomization.yml")
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
