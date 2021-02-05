package config

import "github.com/hashicorp/hcl/v2"

type Component struct {
	Name       string                    `hcl:"name,label"`
	Namespace  string                    `hcl:"namespace"`
	Version    string                    `hcl:"version"`
	RawParams  map[string]*hcl.Attribute `hcl:",remain"`
	JsonParams string
}
