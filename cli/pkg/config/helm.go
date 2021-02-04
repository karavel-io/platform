package config

import (
	helmclient "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/repo"
)

const HelmDefaultRepoName = "karavel"
const HelmDefaultRepo = "https://charts.mikamai.com/karavel"

func (c *Config) SetupHelm() error {
	h, err := helmclient.New(&helmclient.Options{})
	if err != nil {
		return err
	}

	for i := range c.Components {
		c.Components[i].h = h
	}

	return h.AddOrUpdateChartRepo(repo.Entry{
		Name: c.helmRepoName,
		URL:  c.HelmRepoUrl,
	})
}
