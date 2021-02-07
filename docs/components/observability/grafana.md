# Grafana

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries.grafana[0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./grafana.md)
[![Grafana Operator version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=grafana+operator+version&query=$.entries.grafana[0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://example.com)

# Overview

Karavel provides an instance of [Grafana] managed by the [grafana-operator] that is used as the core
observability center for the entire cluster. Grafana aggregates logs, metrics and tracing information so that
they can be easily queried and correlated.

The [grafana-operator] allows to manage Grafana instances, dashboards and datasources as Kubernetes resources.
This allows users to add dashboards to their application deployment manifests and have them automatically
provisioned on Grafana. Check the [official documentation](https://github.com/integr8ly/grafana-operator/tree/master/documentation)
for more information.

## Component configuration

```hcl
component "grafana" {
  version = "0.1.0"
  namespace = "monitoring"

  # Params default values
  
  my = {
    awesome = "param"
  }
}
```

[Grafana]: https://grafana.com/oss/grafana
[grafana-operator]: https://github.com/integr8ly/grafana-operator
[Loki]: https://grafana.com/oss/loki
[Promtail]: https://grafana.com/docs/loki/latest/clients/promtail/
[Explorer tab]: https://grafana.com/docs/grafana/latest/explore/
