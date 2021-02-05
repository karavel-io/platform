package config

import (
	"errors"
	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/gohcl"
	"github.com/hashicorp/hcl/v2/hclparse"
	"github.com/mikamai/karavel/cli/pkg/helmw"
	"github.com/zclconf/go-cty/cty"
	"github.com/zclconf/go-cty/cty/json"
	"io"
	"strings"
)

var (
	ErrConfigParseFailed = errors.New("failed to parse Karavel config")
)

type Config struct {
	Components  []Component `hcl:"component,block"`
	HelmRepoUrl string      `hcl:"charts_repo,optional"`
}

func ReadFrom(logw io.Writer, filename string) (Config, error) {
	var c Config

	p := hclparse.NewParser()

	w := hcl.NewDiagnosticTextWriter(logw, p.Files(), 79, true)
	f, err := p.ParseHCLFile(filename)
	if err != nil {
		_ = w.WriteDiagnostics(err)
		if err.HasErrors() {
			return c, ErrConfigParseFailed
		}
	}

	if err := gohcl.DecodeBody(f.Body, nil, &c); err != nil {
		_ = w.WriteDiagnostics(err)
		if err.HasErrors() {
			return c, ErrConfigParseFailed
		}
	}

	if c.HelmRepoUrl == "" {
		c.HelmRepoUrl = helmw.HelmDefaultRepo
	}

	for i := range c.Components {
		cc := &c.Components[i]
		cc.Name = strings.ToLower(cc.Name)
		pp := make(map[string]cty.Value)
		for l, a := range cc.RawParams {
			v, err := a.Expr.Value(nil)
			if err != nil {
				_ = w.WriteDiagnostics(err)
				if err.HasErrors() {
					return c, ErrConfigParseFailed
				}
			}
			pp[l] = v
		}
		m := cty.ObjectVal(pp)
		j, jerr := json.Marshal(m, m.Type())
		if jerr != nil {
			return c, jerr
		}
		cc.JsonParams = string(j)
	}

	return c, nil
}
