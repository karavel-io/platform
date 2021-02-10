# Karavel Components

Karavel components are software bundles that install and configure applications on Kubernetes.
They range from observability tools like [Grafana] and [Prometheus], to networking addons like [Calico]
and automation tools like [External DNS] and [cert-manager]. These software packages include carefully crafted
Kubernetes manifests for these applications, with hand-picked configurations and integrations between them
to provide a comprehensive and opinionated Kubernetes platform for enterprises to deploy their workloads on.

!!!info
    Components marked with the ![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge) badge
    are essential for bootstrapping the Karavel Platform and allow ArgoCD to kick in and start the GitOps loop.

## Core components

- [Calico] - networking plugin
- [ArgoCD] - GitOps-based Continuous Delivery platform
- [External Secrets] - secrets management system
- [External DNS] - DNS record manager
- [cert-manager] - TLS certificates manager
- [NGINX Ingress Controller] - in-cluster HTTP load balancer

## Security components

Coming Soon

## Observability components

- [Grafana] - dashboard visualization interface
- [Prometheus] - Monitoring system and metrics database
- [Loki] - log collection and query server
- [Tempo] - distributed tracing collector
- [Goldpinger] - cluster nodes connectivity debugger

## AWS components

These components are only needed for clusters running on AWS infrastructure, either regular EC2 or managed EKS clusters.

- [AWS Node Termination Handler] - gracefully handle EC2 instance shutdown within Kubernetes

[Calico]: core/calico.md
[External DNS]: core/external-dns.md
[cert-manager]: core/cert-manager.md
[ArgoCD]: core/argocd.md
[External Secrets]: core/external-secrets.md
[NGINX Ingress Controller]: core/nginx-ingress-controller.md
[Grafana]: observability/grafana.md
[Prometheus]: observability/prometheus.md
[Loki]: observability/loki.md
[Tempo]: observability/tempo.md
[Goldpinger]: observability/goldpinger.md
[AWS Node Termination Handler]: aws/aws-node-termination-handler.md
