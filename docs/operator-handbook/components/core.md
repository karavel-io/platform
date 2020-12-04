# Core components

The [Karavel Bootstrap Tool] includes a `core` role that will provide
core infrastructure addons necessary to all Karavel clusters. These components are fundamental for clusters
to operate properly.

## KBT Import

```yaml
import_role:
  name: mikamai.karavel.core
```

## Components

- [Calico]  
  Network plugin, implements the Container Network Interface and provides
  support for Kubernetes NetworkPolicies  
  [config](./variables.md#calico)
- [External Secrets]  
  Kubernetes controller that provisions regular Kubernetes Secret objects from
  an external vault, like AWS Secrets Manager or Hashicorp Vault  
  [config](./variables.md#external-secrets)
- [External DNS]  
  Kubernetes controller that provisions DNS records based on
  Service and Ingress configurations  
  [config](./variables.md#external-dns)
- [cert-manager]  
  Automates TLS certificates provisioning from various sources, including Let's Encrypt, private CAs and self-signed.   
  [config](./variables.md#cert-manager)
- [Dex]  
  Identity service that uses OpenID Connect to drive authentication for other apps.
  It is used as the central authentication provider for many other services and components,
  delegating to an upstream provider like a company SSO, GitHub and alike  
  [config](./variables.md#dex)
- [ArgoCD]  
  GitOps engine and deployment controller, used to monitor Git repositories containing Kubernetes manifests for changing
  and synchronizing them with the cluster  
  [config](./variables.md#argocd)
- [Goldpinger]  
  Debugging tool for Kubernetes which tests and displays connectivity between nodes in the cluster
  as well as connectivity to the Public Internet
  

[Karavel Bootstrap Tool]: ./bootstrap.md
[Calico]: https://projectcalico.org
[External Secrets]: https://github.com/external-secrets/kubernetes-external-secrets
[External DNS]: https://github.com/kubernetes-sigs/external-dns
[cert-manager]: https://cert-manager.io
[Dex]: https://dexidp.io
[ArgoCD]: https://argoproj.github.io/argo-cd
[Goldpinger]: https://github.com/bloomberg/goldpinger
