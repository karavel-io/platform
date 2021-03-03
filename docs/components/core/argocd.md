# ArgoCD

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries.argocd[0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./argocd.md)
[![ArgoCD version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=argocd+version&query=$.entries.argocd[0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://example.com)
![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge)

## Overview

DESCRIPTION

## Component configuration

```hcl
component "argocd" {
  version = "0.1.0"
  namespace = "argocd"

  # Params default values
  
  my = {
    awesome = "param"
  }
}
```
