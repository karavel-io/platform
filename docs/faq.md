# Frequently Asked Questions

## Karavel Container Platform

### Q: Is the Karavel Container Platform a certified Kubernetes distribution?

**A:** No, we do not bundle Kubernetes nor provide any means of installing it as part of the platform offering. 
The Karavel Container Platform expects a conformant Kubernetes cluster with enough spare capacity to be present before deploying. 
We only provide software components and addons that are deployed on top of the cluster to build up a comprehensive and 
production-grade environment for developers to run and operate their workloads.

The Karavel Container Platform is virtually compatible with any Kubernetes cluster, although there may be some cases where existing addons
may conflict with the Karavel components (for example, Red Hat OpenShift already provides the Prometheus Operator as part of its
core). A bare, upstream Kubernetes without any extra component (not even a CNI) is the ideal starting point for deploying the Platform.

### Q: Are components available as Helm charts?

**A:** Well yes, but actually no. While the platform components are indeed 
[packaged as Helm charts](https://github.com/mikamai/karavel/tree/master/components), they are not meant to be installed
with Helm (e.g. via `helm install`). Instead, they are consumed by the [Karavel CLI], templated based on the provided
Karavel configuration file, and then wrote to disk in a specific directory structure (documented [here](quickstart.md#bootstrap)) that
will be set-up as Kustomize stacks. These Kustomize stacks are then deployed by [ArgoCD] onto the target cluster.  

While you *could* go ahead and install the components via Helm directly, they are not designed for this usage and have way less 
configuration parameters than the upstream chart they are based on. They provide highly opinionated configurations
that are meant to work in concert with the rest of the Karavel stack and any customization is handled through Kustomize patches. 

If you want to install, say, [Prometheus] with Helm, you should use the [official chart](https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack) 
instead.

### Q: Can I swap component X for alternative Y?

**A:** The Karavel Container Platform builds on a set of principles and assumptions based on the team's experience running Kubernetes in production.
The available components have been selected and their configuration carefully crafted to integrate with each other and compose a robust
production-grade platform with a great developer experience out of the box, so while a few of them have alternative implementations that can be swapped
in and out to accomodate specific scenarios, it is unlikely that we'll provide other solutions for the core components.

If you have a specific need that is not satisfied by the platform current state, please reach out to the [maintainers], we'll be happy to help you!

[ArgoCD]: components/core/argocd.md
[Karavel CLI]: cli.md
[Prometheus]: https://prometheus.io
[maintainers]: https://github.com/mikamai/karavel/graphs/contributors
