package helmw

import (
	"bytes"
	"fmt"
	helmclient "github.com/mittwald/go-helm-client"
	"gopkg.in/yaml.v3"
	"helm.sh/helm/v3/pkg/action"
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chart/loader"
)

func GetChartManifest(chartname string) (*chart.Metadata, error) {
	chartname = fmt.Sprintf("%s/%s", HelmRepoName, chartname)
	hc, err := helmclient.New(&helmclient.Options{})
	if err != nil {
		return nil, err
	}
	h := hc.(*helmclient.HelmClient)
	hshow := action.NewShow(action.ShowAll)
	path, err := hshow.ChartPathOptions.LocateChart(chartname, h.Settings)
	if err != nil {
		return nil, err
	}

	ch, err := loader.Load(path)
	if err != nil {
		return nil, err
	}

	return ch.Metadata, nil
}

type YamlDoc map[string]interface{}

func TemplateChart(name string, namespace string, version string) ([]YamlDoc, error) {
	hc, err := helmclient.New(&helmclient.Options{})
	if err != nil {
		return nil, err
	}
	h := hc.(*helmclient.HelmClient)

	ch := &helmclient.ChartSpec{
		ChartName: fmt.Sprintf("%s/%s", HelmRepoName, name),
		Namespace: namespace,
		Version:   version,
		SkipCRDs:  false,
	}

	manifests, err := h.TemplateChart(ch)
	if err != nil {
		return nil, err
	}

	dec := yaml.NewDecoder(bytes.NewReader(manifests))

	var docs []YamlDoc
	for {
		var doc YamlDoc
		if dec.Decode(&doc) != nil {
			break
		}
		docs = append(docs, doc)
	}
	return docs, nil
}
