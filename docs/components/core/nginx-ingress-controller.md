# NGINX Ingress Controller

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries['ingress-nginx'][0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./nginx-ingress-controller.md)
[![NGINX Ingress Controller version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=nginx+ingress+controller+version&query=$.entries['ingress-nginx'][0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://example.com)
![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge)

## Overview

DESCRIPTION

## Component configuration

```hcl
component "ingress-nginx" {
  version = "0.1.0"
  namespace = "ingress-nginx"

  # Params default values
  
  my = {
    awesome = "param"
  }
}
```

[Kubernetes Ingress Controllers]: https://kubernetes.io/docs/concepts/services-networking/ingress-controllers
[ingress class annotation]: https://kubernetes.github.io/ingress-nginx/user-guide/multiple-ingress/
[NGINX]: https://kubernetes.github.io/ingress-nginx
