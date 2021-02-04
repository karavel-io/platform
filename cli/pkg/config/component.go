package config

type Component struct {
	Name      string `hcl:"name,label"`
	Namespace string `hcl:"namespace"`
	Version   string `hcl:"version"`
}
