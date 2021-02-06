package utils

import (
	"bytes"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"modernc.org/sortutil"
	"os"
	"path/filepath"
	"sigs.k8s.io/kustomize/api/types"
	"sort"
)

func RenderKustomizeFile(outdir string, resources []string, ignoreFn func(s string) bool) error {
	filename := filepath.Join(outdir, "kustomization.yml")
	exists := false
	info, err := os.Stat(filename)
	if err == nil || !os.IsNotExist(err) {
		exists = true
	}

	if info != nil && info.IsDir() {
		return errors.Errorf("could not render %s: is a directory", filename)
	}

	var kfile types.Kustomization
	if exists {
		f, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}

		if err := kfile.Unmarshal(f); err != nil {
			return err
		}
		rr := make([]string, 0)
		for _, r := range kfile.Resources {
			if !ignoreFn(r) {
				rr = append(rr, r)
			}
		}

		resources = append(resources, rr...)
		resources = resources[:sortutil.Dedupe(sort.StringSlice(resources))]
	}

	kfile.FixKustomizationPostUnmarshalling()

	kfile.Resources = resources

	kfile.FixKustomizationPreMarshalling()

	var buf bytes.Buffer
	enc := yaml.NewEncoder(&buf)
	enc.SetIndent(2)
	if err := enc.Encode(&kfile); err != nil {
		return err
	}

	if err := enc.Close(); err != nil {
		return err
	}

	if err := ioutil.WriteFile(filename, buf.Bytes(), 0655); err != nil {
		return err
	}

	return nil
}
