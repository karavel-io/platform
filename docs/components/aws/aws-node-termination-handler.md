# AWS Node Termination Handler

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries['aws-node-termination-handler'][0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./aws-node-termination-handler.md)
[![AWS Node Termination Handler version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=aws+node+termination+handler+version&query=$.entries['aws-node-termination-handler'][0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://github.com/aws/aws-node-termination-handler)
![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge)

## Overview

DESCRIPTION

## Component configuration

```hcl
component "aws-node-termination-handler" {
  version = "0.1.0"
  namespace = "kube-system"

  # Params default values
  
  my = {
    awesome = "param"
  }
}
```
