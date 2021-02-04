package utils

import (
	"bytes"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"path/filepath"
	"sigs.k8s.io/kustomize/api/types"
)

func RenderKustomizeFile(outdir string, resources []string) error {
	var kfile types.Kustomization
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

	filename := filepath.Join(outdir, "kustomization.yml")
	if err := ioutil.WriteFile(filename, buf.Bytes(), 0655); err != nil {
		return err
	}

	return nil
}
