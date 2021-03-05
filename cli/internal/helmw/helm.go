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

package helmw

import (
	helmclient "github.com/mittwald/go-helm-client"
	"helm.sh/helm/v3/pkg/repo"
)

const HelmRepoName = "karavel"
const HelmDefaultStableRepo = "https://charts.mikamai.com/karavel"
const HelmDefaultEdgeRepo = "https://charts.mikamai.com/karavel-edge"

func SetupHelm(repoUrl string) error {
	h, err := helmclient.New(&helmclient.Options{})
	if err != nil {
		return err
	}

	if repoUrl == "" {
		repoUrl = HelmDefaultStableRepo
	}

	return h.AddOrUpdateChartRepo(repo.Entry{
		Name: HelmRepoName,
		URL:  repoUrl,
	})
}
