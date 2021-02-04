package helmw

import (
	helmclient "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/repo"
)

const HelmRepoName = "karavel"
const HelmDefaultRepo = "https://charts.mikamai.com/karavel"

func SetupHelm(repoUrl string) error {
	h, err := helmclient.New(&helmclient.Options{})
	if err != nil {
		return err
	}

	if repoUrl == "" {
		repoUrl = HelmDefaultRepo
	}

	return h.AddOrUpdateChartRepo(repo.Entry{
		Name: HelmRepoName,
		URL:  repoUrl,
	})
}
