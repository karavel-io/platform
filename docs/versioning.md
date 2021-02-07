# Versioning

The Karavel Platform aims at supporting the current version of [Kubernetes] plus **two** prior versions. This
provides enough time for operators to maintain their Karavel clusters with confidence and clear upgrade paths.

Karavel maintains two versioning policies: a **per-component policy** and a **platform-wide policy**.
All versions follow [Semantic Versioning] principles and are always in the `major.minor.patch` format.

## Component versioning policy

Each Karavel component is packaged and versioned individually. Component versions may increase because of
changes being made in the component chart, or because the packaged software version has been updated. 

In the latter case the component version will be incremented as well to reflect the new upstream version.  
For instance, if [cert-manager] updates from `1.0.0` to `1.0.1`, the component's **patch** version will be increased.  
If [cert-manager] updates from `1.0.0` to `1.1.0`, this will update the component's **minor** version.  
And if [cert-manager] releases a shiny new 2.0.0 version, a new **major** version of the component will be released as well

Sometimes however, the Karavel project will release a new version of a component to reflect changes made to the chart
or the manifests it deploys, without changing the upstream software version. 
These changes will trigger a version bump according to the rules of [Semantic Versioning].  
A backward compatible bugfix will increment the **patch** version.  
A backward compatible improvement or new feature will increment the **minor** version.  
A backward **incompatible** or otherwise breaking change will increment the **major** version.  

## Platform versioning policy

Due to the large amount of components provided by Karavel and the possible interactions between them to compose
a production-ready Kubernetes platform, the Karavel project regularly picks specific component versions that have been
extensively tested and verified to work well together. These curated selections compose the **Karavel Platform** and are published
under a single version, much like a Linux distribution would do with its packages and repositories. This gives developers working
on components the freedom to iterate and improve their charts without being tied down to the others, and it gives operators
the confidence that their clusters will always work at their best, while still being able to update specific components independently
to leverage new features.

The Karavel team maintains the Karavel Platform in sync with the target [Kubernetes] versions.  
Karavel supports the **current** K8s version plus **two** prior releases.  
The Karavel project will drop support for a given Karavel Platform release when the lowest compatible Kubernetes version
becomes unsupported. As a general recommendation, operators should strive to keep Karavel clusters on the latest minor
version of both Kubernetes and the Karavel Platform.

### Kubernetes compatibility matrix

| Karavel Platform Version  | Kubernetes 1.18    | Kubernetes 1.19    | Kubernetes 1.20 (current)    |
|:--------------------------|:------------------:|:------------------:|:----------------------------:|
| 0.1.0                     | :material-close:   | :material-alert:   | :material-check:             |

*:material-check: compatible*  
*:material-alert: specific requirements or caveats are present. Consult the release changelog for more information*  
*:material-close: incompatible or unsupported*  

[Kubernetes]: https://kubernetes.io/docs/setup/release/version-skew-policy/
[Semantic Versioning]: https://semver.org
[cert-manager]: components/core/cert-manager.md
