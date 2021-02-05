package plan

import (
	"fmt"
	"github.com/mikamai/karavel/cli/pkg/config"
	"github.com/mikamai/karavel/cli/pkg/helmw"
	"github.com/pkg/errors"
)

type Plan struct {
	components map[string]*Component
}

func NewFromConfig(cfg *config.Config) (*Plan, error) {
	p := New()

	for _, c := range cfg.Components {
		meta, err := helmw.GetChartManifest(c.Name)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build plan from config")
		}
		comp, err := NewComponentFromChartMetadata(meta)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build plan from config")
		}
		comp.namespace = c.Namespace
		comp.jsonParams = c.JsonParams
		p.AddComponent(comp)
	}

	return &p, nil
}

func New() Plan {
	return Plan{
		components: map[string]*Component{},
	}
}

func (p *Plan) Components() []*Component {
	var cc []*Component
	for _, c := range p.components {
		cc = append(cc, c)
	}

	return cc
}

func (p *Plan) GetComponent(name string) *Component {
	return p.components[name]
}

func (p *Plan) AddComponent(c Component) {
	p.components[c.name] = &c
}

func (p *Plan) HasComponent(name string) bool {
	return p.components[name] != nil
}

func (p *Plan) Validate() error {
	if err := p.checkDependencies(); err != nil {
		return err
	}

	if err := p.processIntegrations(); err != nil {
		return err
	}

	return nil
}

func (p *Plan) checkDependencies() error {
	for n, c := range p.components {
		for _, dn := range c.dependencies {
			if !p.HasComponent(dn) {
				return fmt.Errorf("missing dependency: component '%s' requires '%s'", n, dn)
			}
		}
	}
	return nil
}

func (p *Plan) processIntegrations() error {
	for _, c := range p.components {
		c.integrations = make(map[string]bool)
		for integ, dd := range c.integrationsDeps {
			active := len(dd) > 0
			for _, dn := range dd {
				active = active && p.HasComponent(dn)
			}
			c.integrations[integ] = active
		}
		if err := c.patchIntegrations(); err != nil {
			return err
		}
	}
	return nil
}
