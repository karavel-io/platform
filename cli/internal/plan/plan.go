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

package plan

import (
	"fmt"
	"github.com/mikamai/karavel/cli/internal/helmw"
	"github.com/mikamai/karavel/cli/pkg/config"
	"github.com/pkg/errors"
)

type Plan struct {
	components map[string]*Component
}

func NewFromConfig(cfg *config.Config) (*Plan, error) {
	p := New()

	for _, c := range cfg.Components {
		chartName := c.Name
		if c.ComponentName != "" {
			chartName = c.ComponentName
		}

		meta, err := helmw.GetChartManifest(chartName)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build plan from config")
		}
		comp, err := NewComponentFromChartMetadata(meta)
		if err != nil {
			return nil, errors.Wrap(err, "failed to build plan from config")
		}
		if c.ComponentName != "" {
			comp.name = c.Name
		}
		comp.namespace = c.Namespace
		comp.jsonParams = c.JsonParams
		if err := p.AddComponent(comp); err != nil {
			return nil, err
		}
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

func (p *Plan) AddComponent(c Component) error {
	if p.components[c.name] != nil {
		return errors.Errorf("duplicate component '%s' found", c.name)
	}
	p.components[c.name] = &c
	return nil
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
