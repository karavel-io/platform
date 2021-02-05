package plan

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/argo"
	"github.com/mikamai/karavel/cli/pkg/helmw"
	"github.com/mikamai/karavel/cli/pkg/utils"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/chart"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"sync"
)

// reserved annotations that should not be interpreted as
// feature flags
var reservedAnnotations = map[string]bool{
	"bootstrap": true,
}

type Component struct {
	name             string
	namespace        string
	version          string
	bootstrap        bool
	dependencies     []string
	integrationsDeps map[string][]string
	integrations     map[string]bool
}

func NewComponentFromChartMetadata(meta *chart.Metadata) (Component, error) {
	var deps []string
	for _, dep := range meta.Dependencies {
		deps = append(deps, dep.Name)
	}

	integs := make(map[string][]string)
	for integ, s := range meta.Annotations {
		if reservedAnnotations[integ] {
			continue
		}

		r := strings.NewReader(s)
		cr := csv.NewReader(r)
		adeps, err := cr.Read()
		if err != nil {
			return Component{}, errors.Wrap(err, "failed to read integration dependencies")
		}
		integs[integ] = adeps
	}

	bootstrap, err := strconv.ParseBool(meta.Annotations["bootstrap"])
	if err != nil {
		bootstrap = false
	}

	return Component{
		name:             meta.Name,
		version:          meta.Version,
		bootstrap:        bootstrap,
		dependencies:     deps,
		integrationsDeps: integs,
	}, nil
}

func (c *Component) Name() string {
	return c.name
}

func (c *Component) Namespace() string {
	return c.namespace
}

func (c *Component) Version() string {
	return c.version
}

func (c *Component) IsBootstrap() bool {
	return c.bootstrap
}

type routineRes struct {
	filename string
	err      error
}

func (c *Component) Render(outdir string) error {
	deferr := fmt.Sprintf("failed to render component '%s' v%s", c.name, c.version)

	if err := os.RemoveAll(outdir); err != nil {
		return errors.Wrap(err, deferr)
	}

	if err := os.MkdirAll(outdir, 0755); err != nil {
		return errors.Wrap(err, deferr)
	}

	docs, err := helmw.TemplateChart(c.name, c.namespace, c.version)
	if err != nil {
		return errors.Wrap(err, deferr)
	}

	var wg sync.WaitGroup
	resch := make(chan routineRes)

	for i, d := range docs {
		wg.Add(1)
		go func(i int, doc helmw.YamlDoc) {
			defer wg.Done()
			if len(doc) == 0 {
				return
			}

			var buf bytes.Buffer
			enc := yaml.NewEncoder(&buf)
			enc.SetIndent(2)
			if err := enc.Encode(&doc); err != nil {
				resch <- routineRes{err: err}
				return
			}

			if err := enc.Close(); err != nil {
				resch <- routineRes{err: err}
				return
			}

			k := strings.ToLower(doc["kind"].(string))
			meta := doc["metadata"].(helmw.YamlDoc)
			ns := ""
			if meta["namespace"] != nil && meta["namespace"] != c.namespace {
				ns = "-" + meta["namespace"].(string)
			}

			basename := fmt.Sprintf("%s-%s%s.yml", meta["name"], k, ns)
			filename := filepath.Join(outdir, basename)
			if err := ioutil.WriteFile(filename, buf.Bytes(), 0655); err != nil {
				resch <- routineRes{err: err}
			}
			resch <- routineRes{filename: filepath.Base(basename)}
		}(i, d)
	}

	go func() {
		wg.Wait()
		close(resch)
	}()

	var files []string
	for res := range resch {
		if res.err != nil {
			return errors.Wrap(res.err, deferr)
		}
		files = append(files, res.filename)
	}

	return utils.RenderKustomizeFile(outdir, files)
}

func (c *Component) RenderApplication(argoNs string, repoUrl string, path string, outfile string) error {
	deferr := fmt.Sprintf("failed to render application manifest for component '%s' v%s", c.name, c.version)

	app := argo.NewApplication(c.name, c.namespace, argoNs, repoUrl, path)

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(&app); err != nil {
		return errors.Wrap(err, deferr)
	}

	if err := enc.Close(); err != nil {
		return errors.Wrap(err, deferr)
	}

	if err := ioutil.WriteFile(outfile, buf.Bytes(), 0655); err != nil {
		return errors.Wrap(err, deferr)
	}
	return nil
}