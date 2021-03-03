## What is Karavel?

Karavel is a project that provides tools and knowledge around the [Kubernetes]
stack to deploy and manage [private PaaS] in the Cloud and on premise.

The main output of the Karavel project is the [Karavel Platform](#the-karavel-platform), a curated
set of components and services based on the best-in-breed open source projects, carefully configured to deliver
a production-ready platform for enterprises.

## The Karavel Platform

Karavel provides many different [components] that together form a cohesive and integrated environment
called the **Karavel Platform**.

The Karavel Platform selects specific versions of these components that are tried and tested together
and publishes them as a ready-to-use kit to assemble GitOps-enabled platforms for enterprises. Check out
our [quickstart guide] for an introductory view to the Platform.

### Features

- Run on any [conformant Kubernetes cluster]
- 100% open source stack based on community and CNCF projects
- GitOps first workflow, enabling the platform to be self-hosted and updating itself
- Built-in security tools for secrets management, policy enforcement and access control
- Elastic routing layer with automated DNS, load balancing and certificate management
- Comprehensive observability stack with metrics, logging and distributed tracing collection and visualization

[Kubernetes]: https://kubernetes.io
[private PaaS]: https://en.wikipedia.org/wiki/Platform_as_a_service#Public,_private_and_hybrid
[components]: ./components/index.md
[quickstart guide]: ./quickstart.md
[conformant Kubernetes cluster]: https://www.cncf.io/certification/software-conformance/
