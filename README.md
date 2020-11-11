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
> that are designed to be installed in a certain way. For instance, they do not honor 
> Helm's release namespace and instead manage their own Kubernetes namespaces.
> If you need more flexibility regarding how each component is installed you should
> manage it separately using the recommended upstream installation method.

## Requirements

Since MKP is distributed in the form of Helm Charts it can be installed by
simply running `helm install` for each module. However we recommend the use of Helmfile
to declaratively install all the required packages in one command.   
The following CLI tools are required:

- [helm] to install the charts
- [helmfile] to orchestrate chart installation
- [kubectl], used by helm to communicate with the Kubernetes cluster.

## Quickstart

MKP is built on the principles of [GitOps], and as such is designed to
keep the desired state of the platform in a git repository.

```bash
mkdir mkp && cd mkp
git init
```

Create a `helmfile.yaml` to list the desired modules.
MKP is designed to be modular, so that some components can be replaced or omitted altogether.
Core components are always required.

```yaml
repositories:
  - name: mkp
    url: https://mkp.charts.mikamai.com

releases:
  - name: core
    chart: mkp/core
    version: 0.1.0
  - name: nginx
    chart: mkp/ingress-nginx
    version: 0.1.0
  - name: argocd
    chart: mkp/argocd
    version: 0.1.0
  - name: metrics
    chart: mkp/metrics-prometheus
    version: 0.1.0
  - name: logging
    chart: mkp/logging-loki
    version: 0.1.0
  - name: tracing
    chart: mkp/tracing-tempo
    version: 0.1.0
```

To install them, first lock the chart versions then apply the stack.

```bash
helmfile deps
helmfile apply
```

Finally, check that everything installed correctly.
```bash
kubectl get pods --all-namespaces
```

You can also run chart tests if you want.
```bash
helmfile test --cleanup
```

[Helm Charts]: https://helm.sh
[helm]: https://helm.sh/docs/intro/install/
[helmfile]: https://github.com/roboll/helmfile
[kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl
[GitOps]: https://www.weave.works/blog/what-is-gitops-really

