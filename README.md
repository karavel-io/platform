# Karavel <img align="right" width=384 src="https://via.placeholder.com/384x200.png?text=Karavel%20Logo%20Here">
[![Karavel](https://circleci.com/gh/mikamai/karavel/tree/master.svg?style=svg)](https://circleci.com/gh/mikamai/karavel/tree/master)

Karavel is an pure-upstream Kubernetes distribution
that combines best-in-breed open source components to deliver
a production ready platform for enterprises based on GitOps.

Karavel is packaged as a curated set of components and services, ranging from networking addons
to load balancers and observability tools.

It can be installed on any Kubernetes cluster and once bootstrapped is self-healing and self-hosted, meaning it is capable
of updating itself.

## Available components

- [Calico]
- [ArgoCD]
- [Dex]
- [cert-manager]
- [ExternalDNS]
- [External Secrets]
- [Goldpinger]
- [NGINX Ingress Controller]

## Requirements

TBD

## Quickstart

Karavel uses [Ansible] as its bootstrapping tool. A quickstart guide for setting
up a new Karavel repository can be found in the [examples folder].

[Calico]: https://projectcalico.org
[ArgoCD]: https://argoproj.github.io/argo-cd
[Dex]: https://dexidp.io
[cert-manager]: https://cert-manager.io
[ExternalDNS]: https://github.com/kubernetes-sigs/external-dns
[External Secrets]: https://github.com/godaddy/kubernetes-external-secrets
[Goldpinger]: https://github.com/bloomberg/goldpinger
[NGINX Ingress Controller]: https://kubernetes.github.io/ingress-nginx/
[helm]: https://helm.sh/docs/intro/install/
[helmfile]: https://github.com/roboll/helmfile
[kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl
[GitOps]: https://www.weave.works/blog/what-is-gitops-really
[Ansible]: https://ansible.com
[examples folder]: ./examples/README.md
