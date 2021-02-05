package action

import (
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/config"
	"github.com/mikamai/karavel/cli/pkg/helmw"
	"github.com/mikamai/karavel/cli/pkg/plan"
	"github.com/mikamai/karavel/cli/pkg/utils"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"sync"
)

type RenderParams struct {
	ConfigPath string
	Debug      bool
}

func Render(logger *log.Logger, params RenderParams) error {
	debug := params.Debug
	cpath := params.ConfigPath
	workdir := filepath.Dir(cpath)
	appsDir := filepath.Join(workdir, "applications")
	projsDir := filepath.Join(workdir, "projects")

	logger.Printf("Rendering new Karavel project with config file %s\n", cpath)

	if debug {
		logger.Print("Reading config file")
	}
	cfg, err := config.ReadFrom(logger.Writer(), cpath)
	if err != nil {
		return errors.Wrap(err, "failed to read config file")
	}

	if debug {
		logger.Printf("Setting up Karavel Charts repository %s", cfg.HelmRepoUrl)
	}
	if err := helmw.SetupHelm(cfg.HelmRepoUrl); err != nil {
		return errors.Wrap(err, "failed to setup Karavel Charts repository")
	}

	logger.Println()

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
	for _, c := range p.Components() {
		appFile := c.Name() + ".yml"
		apps = append(apps, appFile)
		if c.IsBootstrap() {
			renderDirs = append(renderDirs, filepath.Join("vendor", c.Name()))
		}

		wg.Add(1)
		go func(comp *plan.Component) {
			defer wg.Done()

			outdir := filepath.Join(workdir, "vendor", comp.Name())
			logger.Printf("Rendering component '%s' %s at %s", comp.Name(), comp.Version(), strings.ReplaceAll(outdir, filepath.Dir(workdir)+"/", ""))
			if debug {
				logger.Printf("Component '%s' %s params: %s", comp.Name(), comp.Version(), comp.Params())
			}
			if err := comp.Render(outdir); err != nil {
				pair, _ := utils.NewPair(comp.Name(), err)
				ch <- pair
				return
			}

			if debug {
				logger.Printf("Rendering application manifest for component '%s' %s", comp.Name(), comp.Version())
			}
			// TODO: git integration to detect repo and path if not provided in config
			if err := comp.RenderApplication(argoNs, "TODO", "TODO", filepath.Join(appsDir, appFile)); err != nil {
				pair, _ := utils.NewPair(comp.Name(), err)
				ch <- pair
			}
		}(c)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for pair := range ch {
		err := pair.ErrorB()
		if err != nil {
			return errors.Wrapf(err, "failed to render component '%s'", pair.A())
		}
	}

	sort.Strings(apps)
	if err := utils.RenderKustomizeFile(appsDir, apps); err != nil {
		return errors.Wrap(err, "failed to render applications kustomization.yml")
	}

	infraProj := "infrastructure.yml"
	if err := ioutil.WriteFile(filepath.Join(projsDir, infraProj), []byte(fmt.Sprintf(argoProject, argoNs)), 0655); err != nil {
		return errors.Wrap(err, "failed to render infrastructure project file")
	}

	if err := utils.RenderKustomizeFile(projsDir, []string{infraProj}); err != nil {
		return errors.Wrap(err, "failed to render projects kustomization.yml")
	}

	if err := utils.RenderKustomizeFile(workdir, renderDirs); err != nil {
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
