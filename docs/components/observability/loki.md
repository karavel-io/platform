# Loki

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries.loki[0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./loki.md)
[![Grafana Loki version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=grafana+loki+version&query=$.entries.loki[0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://example.com)

# Overview

Karavel provides a robust logging stack based on open source projects released by [Grafana Labs].
The log aggregation and query pipeline is provided by [Loki], a very lightweight and powerful 
server developed by Grafana. Logs are collected by a [Promtail] agent installed on each Kubernetes node
and are then accessible and queryable via the [Explorer tab] in Grafana.

## Component configuration

```hcl
component "loki" {
  version = "0.1.0"
  namespace = "monitoring"

  # Params default values
  
  my = {
    awesome = "param"
  }
}
```

[Grafana Labs]: https://grafana.com/oss/
[Grafana]: grafana.md
[Loki]: https://grafana.com/oss/loki
[Promtail]: https://grafana.com/docs/loki/latest/clients/promtail/
[Explorer tab]: https://grafana.com/docs/grafana/latest/explore/
