# Calico

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries.calico[0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./calico.md)
[![calico version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=calico+version&query=$.entries.calico[0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://projectcalico.org)
![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge)

## Overview

[Calico] is an open source networking and network security solution for containers, virtual machines, and native host-based workloads.
It features a Kubernetes-compatible [CNI plugin] that provide the necessary networking functionalities for a cluster to operate.
For example, it assigns interfaces and IPs to pods so that they can talk to each other and to the Kubernetes API.

In addition to the core networking functionalities, Calico also implements the [Kubernetes Network Policies] API that is 
used to partition the otherwise flat cluster network and restrict how a pod can talk to the rest of the network.

!!!Warning
    Calico provides an overlay network with in-cluster private IPs. When deployed to cloud VPCs such as AWS, it
    is often more useful or required that pod IPs are reachable by targets outside the cluster, such as load balancers
    and other private services. To accomodate these use cases, it is possible to deploy Calico without the CNI plugin
    to provide Network Policies functionalities, and delegate the CNI provisioning to another plugin, such as the 
    [AWS VPC CNI plugin]. We are working on a possible alternative solution to better accomodate these cases in the future.

## Component configuration

```hcl
component "calico" {
  version = "0.1.0"
  namespace = "kube-system"

  # Params default values

  cni = {
    # Install Calico CNI plugin. Disable to only use Calico for network policies
    # (e.g. when using AWS VPC CNI plugin)
    enable = true
  }
  typha = {
    # Enable Typha for scaling Calico on big clusters
    enable = false
  }

}
```

[Calico]: https://projectcalico.org
[CNI plugin]: https://github.com/containernetworking/cni
[Kubernetes Network Policies]: https://kubernetes.io/docs/concepts/services-networking/network-policies/
[AWS VPC CNI plugin]: https://docs.aws.amazon.com/eks/latest/userguide/pod-networking.html
