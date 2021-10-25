# Requirements

Karavel is a complex system comprised of a wide selection for components. While the vast majority of them
is entirely self-contained, some pieces need external supporting infrastructure in order to function. This infrastructure
needs to be provisioned beforehand and configured into Karavel.

The documentation for each component lists the exact requirements for each case, but as a general
overview the needed infrastructure parts are:

- a **conformant [Kubernetes] cluster**, of course
- a **secure secrets store** to store credentials and other passwords, like [Hashicorp Vault] or [AWS Secrets Manager],
see [components/external-secrets]
- an S3-compatible **object storage server** to store data like metrics, logs and traces
- a supported DNS provider, see [components/external-dns]

The cluster should be able to host the platform component with enough capacity to spare to enable incremental rollouts
and autoscaling depending on resource usage, in addition to your own workloads. We recommend to start with at least three
nodes with 4 cores and 8 GiB of RAM each, although this can be reduced by using only a few KCP components and scaling down
some of the less critical pods.

[Kubernetes]: https://kubernetes.io
[Hashicorp Vault]: https://vaultproject.io/
[AWS Secrets Manager]: https://aws.amazon.com/secrets-manager/
[components/external-secrets]: /components/external-secrets
[components/external-dns]: /components/external-dns
