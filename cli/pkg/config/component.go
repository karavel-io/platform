package config

import (
	"bytes"
	"fmt"
	helmclient "github.com/mittwald/go-helm-client"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sigs.k8s.io/kustomize/api/types"
	"strings"
	"sync"
)

type Component struct {
	Name      string `hcl:"name,label"`
	Namespace string `hcl:"namespace,optional"`
	Version   string `hcl:"version"`
	h         helmclient.Client
	outdir    string
}

type routineRes struct {
	filename string
	err      error
}

func (c *Component) Render(logger *log.Logger) error {
	deferr := fmt.Sprintf("failed to render component '%s' v%s", c.Name, c.Version)

	if err := os.RemoveAll(c.outdir); err != nil {
		return errors.Wrap(err, deferr)
	}

	if err := os.MkdirAll(c.outdir, 0755); err != nil {
		return errors.Wrap(err, deferr)
	}

	chart := &helmclient.ChartSpec{
		ReleaseName:  c.Name,
		ChartName:    fmt.Sprintf("karavel/%s", c.Name),
		Namespace:    c.Namespace,
		Version:      c.Version,
		NameTemplate: "",
		SkipCRDs:     false,
	}

	manifests, err := c.h.TemplateChart(chart)
	if err != nil {
		return errors.Wrap(err, deferr)
	}

	dec := yaml.NewDecoder(bytes.NewReader(manifests))

	var wg sync.WaitGroup
	resch := make(chan routineRes)

	for {
		var doc map[string]interface{}
		if dec.Decode(&doc) != nil {
			break
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			k := strings.ToLower(doc["kind"].(string))
			meta := doc["metadata"].(map[string]interface{})
			ns := ""
			if meta["namespace"] != nil && meta["namespace"] != c.Namespace {
				ns = "-" + meta["namespace"].(string)
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

			basename := fmt.Sprintf("%s-%s%s.yml", meta["name"], k, ns)
			filename := filepath.Join(c.outdir, basename)
			if err := ioutil.WriteFile(filename, buf.Bytes(), 0655); err != nil {
				resch <- routineRes{err: err}
			}
			resch <- routineRes{filename: filepath.Base(basename)}
		}()
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

	var kfile types.Kustomization
	kfile.FixKustomizationPostUnmarshalling()

	kfile.Resources = files

	kfile.FixKustomizationPreMarshalling()

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(&kfile); err != nil {
		return errors.Wrap(err, deferr)
	}

	if err := enc.Close(); err != nil {
		return errors.Wrap(err, deferr)
	}

	filename := filepath.Join(c.outdir, "kustomization.yml")
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0655); err != nil {
		return errors.Wrap(err, deferr)
	}
	return nil
}
