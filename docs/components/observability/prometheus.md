# Prometheus

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries.prometheus[0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./prometheus.md)
[![Prometheus Operator version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=prometheus+operator+version&query=$.entries.prometheus[0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://example.com)

## Overview

DESCRIPTION

## Component configuration

```hcl
component "prometheus" {
  version = "0.1.0"
  namespace = "monitoring"

  # Params default values
  
  my = {
    awesome = "param"
  }
}
```
