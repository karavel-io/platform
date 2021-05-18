# Karavel <img align="right" width=384 src="https://via.placeholder.com/384x200.png?text=Karavel%20Logo%20Here">

## What is Karavel?

Karavel is a project that provides tools and knowledge around the [Kubernetes]
stack to deploy and manage [private Containers-as-a-Service platforms] in the Cloud and on premise.

The main output of the Karavel project is the [Karavel Container Platform](#the-karavel-platform), a curated
set of components and services based on the best-in-breed open source projects, carefully configured to deliver
a production-ready platform for enterprises.

## The Karavel Container Platform

Karavel provides many different [components] that together form a cohesive and integrated environment
called the **Karavel Container Platform**.

The Karavel Container Platform selects specific versions of these components that are tried and tested together
and publishes them as a ready-to-use kit to assemble GitOps-enabled platforms for enterprises. Check out
our [quickstart guide] for an introductory view to the Platform.

### Features

- Run on any [conformant Kubernetes cluster]
- 100% open source stack based on community and CNCF projects
- GitOps first workflow, enabling the platform to be self-hosted and updating itself
- Built-in security tools for secrets management, policy enforcement and access control
- Elastic routing layer with automated DNS, load balancing and certificate management
- Comprehensive observability stack with metrics, logging and distributed tracing collection and visualization

## Components

The full list of available components is available [here](https://docs.karavel.io/components/).

These are the most notable components offered by the Karavel Container Platform:

- [ArgoCD]
- [Calico]
- [cert-manager]
- [Dex]
- [External DNS]
- [External Secrets]
- [NGINX Ingress Controller]
- [Grafana]
- [Loki]
- [Prometheus]
- [Tempo]

## Quickstart

You can get up and running quickly and efficiently with our [Quickstart Guide].

[Kubernetes]: https://kubernetes.io
[private Containers-as-a-Service platforms]: https://www.redhat.com/en/topics/cloud-computing/what-is-caas
[components]: https://docs.karavel.io/components
[quickstart guide]: https://docs.karavel.io/quickstart
[conformant Kubernetes cluster]: https://www.cncf.io/certification/software-conformance/
[ArgoCD]: https://docs.karavel.io/components/core/argocd
[Calico]: https://docs.karavel.io/components/core/calico
[cert-manager]: https://docs.karavel.io/components/core/cert-manager
[Dex]: https://docs.karavel.io/components/core/dex
[External DNS]: https://docs.karavel.io/components/core/external-dns
[External Secrets]: https://docs.karavel.io/components/core/external-secrets
[NGINX Ingress Controller]: https://docs.karavel.io/components/core/nginx-ingress-controller
[Grafana]: https://docs.karavel.io/components/observability/grafana
[Loki]: https://docs.karavel.io/components/observability/loki
[Prometheus]: https://docs.karavel.io/components/observability/prometheus
[Tempo]: https://docs.karavel.io/components/observability/tempo
