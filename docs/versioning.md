# Versioning

The Karavel Container Platform aims at supporting the current version of [Kubernetes] plus **two** prior versions. This
provides enough time for operators to maintain their Karavel clusters with confidence and clear upgrade paths.

Karavel maintains two versioning policies: a **per-component policy** and a **platform-wide policy**.
Component versions follow [Semantic Versioning] principles and are always in the `major.minor.patch` format (e.g. `0.3.1`).  
Platform versions follow a [Calendar Versioning] schema composed as `yyyy.minor.patch` (e.g. `2021.2.3`).

## Component versioning policy

Each Karavel component is packaged and versioned individually. Component versions may increase because of
changes being made in the component chart, or because the packaged software version has been updated. 

In the latter case the component version will be incremented as well to reflect the new upstream version.  
For instance, if [cert-manager] updates from `1.0.0` to `1.0.1`, the component's **patch** version will be increased.  
If [cert-manager] updates from `1.0.0` to `1.1.0`, this will update the component's **minor** version.  
Finally, if [cert-manager] releases a shiny new 2.0.0 version a new **major** version of the component will be released as well

Sometimes however, the Karavel project will release a new version of a component to reflect changes made to the chart
or the manifests it deploys, without changing the upstream software version. 
These changes will trigger a version bump according to the rules of [Semantic Versioning].  
A backward compatible bugfix will increment the **patch** version.  
A backward compatible improvement or new feature will increment the **minor** version.  
A backward **incompatible** or otherwise breaking change will increment the **major** version.  

## Platform versioning policy

Due to the large amount of components provided by Karavel and the possible interactions between them to compose
a production-ready Kubernetes platform, the Karavel project regularly picks specific component versions that have been
extensively tested and verified to work well together. These curated selections compose the **Karavel Container Platform** and are published
under a single version, much like a Linux distribution would do with its packages and repositories. This gives developers working
on components the freedom to iterate and improve their charts without being tied down to the others, and it gives operators
the confidence that their clusters will always work at their best, while still being able to update specific components independently
to leverage new features.

The Karavel team maintains the Karavel Container Platform in sync with the target [Kubernetes] versions.  
Karavel supports the **current** K8s version plus **two** prior releases.  
The Karavel project will drop support for a given Karavel Container Platform release when the lowest compatible Kubernetes version
becomes unsupported. As a general recommendation, operators should strive to keep Karavel clusters on the latest minor
version of both Kubernetes and the Karavel Container Platform.

### Kubernetes compatibility matrix

| Karavel Container Platform Version  | Kubernetes 1.20       | Kubernetes 1.21       | Kubernetes 1.22 (current) | Kubernetes 1.23 (next)    |
| :---------------------------------- | :-------------------: | :-------------------: | :-----------------------: | :-----------------------: |
| unstable                            | :material-check-bold: | :material-check-bold: | :material-clock:          | :material-clock:          |
| 2021.1 (coming December 2021)       |                       | :material-clock:      | :material-clock:          | :material-clock:          |

*:material-check-bold: compatible*  
*:material-alert: specific requirements or caveats are present. Consult the release changelog for more information*  
*:material-clock: planned but not yet available*  
*:material-wrench: actively being developed*  

### Cloud providers support matrix

The Karavel Container Platform is regularly tested [on a number of different Kubernetes managed services](https://github.com/karavel-io/platform-e2e) in addition
to plain upstream Kubernetes. Currently, these are the officially supported providers. Kubernetes versions are tested following the table in the previous section.
Support for a managed service illustrates the integration status with other parts of the cloud provider, such as their credentials
storage for [External Secrets](/components/external-secrets), [DNS provider](/components/external-dns), object storage, and so on.

| Provider                  | Kubernetes            | Secrets Storage        | DNS                   | Object Storage        |
| :------------------------ | :-------------------: | :--------------------: | :-------------------: | :-------------------: |
| [Amazon EKS]              | :material-check-bold: | :material-check-bold:  | :material-check-bold: | :material-check-bold: |
| [DigitalOcean Kubernetes] | :material-wrench:     | offering not available | :material-wrench:     | :material-check-bold: |
| [Azure AKS]               | :material-clock:      | :material-clock:       | :material-clock:      | :material-clock:      |
| [Google GKE]              | :material-clock:      | :material-clock:       | :material-clock:      | :material-clock:      |

For a more updated timeline of planned features, please refer to the [official project board].

[Kubernetes]: https://kubernetes.io/docs/setup/release/version-skew-policy/
[Semantic Versioning]: https://semver.org
[Calendar Versioning]: https://calver.org
[cert-manager]: components/cert-manager
[Amazon EKS]: https://aws.amazon.com/eks/
[DigitalOcean Kubernetes]: https://www.digitalocean.com/products/kubernetes/
[Azure AKS]: https://azure.microsoft.com/en-us/services/kubernetes-service/
[Google GKE]: https://cloud.google.com/kubernetes-engine
[official project board]: https://github.com/orgs/karavel-io/projects/2
