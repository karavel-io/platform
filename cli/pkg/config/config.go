package config

import (
	"errors"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"io"
	"path/filepath"
	"strings"
)

var (
	ErrConfigParseFailed = errors.New("failed to parse Karavel config")
)

type Config struct {
	Components   []Component `hcl:"component,block"`
	HelmRepoUrl  string      `hcl:"charts_repo,optional"`
	helmRepoName string
	workdir      string
}

func ReadFrom(logw io.Writer, filename string) (Config, error) {
	var c Config

	p := hclparse.NewParser()
	w := hcl.NewDiagnosticTextWriter(logw, p.Files(), 79, true)
	f, err := p.ParseHCLFile(filename)
	if err != nil {
		_ = w.WriteDiagnostics(err)
		return c, ErrConfigParseFailed
	}

	if err := gohcl.DecodeBody(f.Body, nil, &c); err != nil {
		_ = w.WriteDiagnostics(err)
		return c, ErrConfigParseFailed
	}

	c.workdir = filepath.Dir(filename)
	if c.HelmRepoUrl == "" {
		c.HelmRepoUrl = HelmDefaultRepo
	}
	c.helmRepoName = HelmDefaultRepoName

	for i := range c.Components {
		cc := &c.Components[i]
		cc.Name = strings.ToLower(cc.Name)
		cc.outdir = filepath.Join(c.workdir, "vendor", cc.Name)
	}

	return c, nil
}
