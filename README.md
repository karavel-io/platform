# Mikamai Kubernetes Platform

MKP is an pure-upstream Kubernetes distribution
that combines best-in-breed components to deliver
a production ready platform for enterprises.

MKP is packaged as a collection of [Helm Charts]
that install a curated set of components and services, ranging from networking addons
to load balancers and observability tools.

It can be installed on any Kubernetes cluster and once bootstrapped is self-healing and
capable of managing itself.

> :warning: MKP charts are highly opinionated software bundles
> that are designed to be installed in a certain way. For instance, they do not always honor
> Helm's release namespace and instead manage their own Kubernetes namespaces.
> Also they provide a very small set of configurable values compared to the upstream charts they are
> based of.
> If you need more flexibility regarding how each component is installed you should
> manage it separately using the recommended upstream installation method.

## Available modules

- [ArgoCD]
- [Calico]
- [cert-manager]
- [ExternalDNS]
- [External Secrets]
- [Goldpinger]
- [NGINX Ingress Controller]

## Requirements

Since MKP is distributed in the form of Helm Charts it can be installed by
simply running `helm install` for each module. However we recommend the use of Helmfile
to declaratively install all the required packages in one command.
The following CLI tools are required:

- [helm] to install the charts
- [helmfile] to orchestrate chart installation and generate the GitOps solution
- [kubectl] used to interact with the cluster

## Quickstart

MKP is built on the principles of [GitOps], and as such is designed to
keep the desired state of the platform in a git repository.

Check the following examples for more information about setting up your MKP cluster:

- [EKS Quickstart]

[Helm Charts]: https://helm.sh
[ArgoCD]: https://argoproj.github.io/argo-cd
[Calico]: https://projectcalico.org
[cert-manager]: https://cert-manager.io
[ExternalDNS]: https://github.com/kubernetes-sigs/external-dns
[External Secrets]: https://github.com/godaddy/kubernetes-external-secrets
[Goldpinger]: https://github.com/bloomberg/goldpinger
[NGINX Ingress Controller]: https://kubernetes.github.io/ingress-nginx/
[helm]: https://helm.sh/docs/intro/install/
[helmfile]: https://github.com/roboll/helmfile
[kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl
[GitOps]: https://www.weave.works/blog/what-is-gitops-really
[EKS Quickstart]: ./examples/eks/README.md
