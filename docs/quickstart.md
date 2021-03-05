# Quickstart

This section will guide you through the first bootstrap
of a new Karavel cluster from scratch. It assumes you have a general knowledge
of Git and Kubernetes and know your way around a command shell.

## Requirements

### Tools

- the [karavel] CLI tool
- [git], [kubectl], [kustomize] installed on your local machine
- admin access to a running Kubernetes cluster. Minikube is fine for a simple, local deployment for testing purposes.

The `karavel` CLI can be downloaded from [GitHub](https://github.com/mikamai/karavel/releases).  
Alternatively, you can build it from source. You need Golang 1.15+ installed to build it.

`go get -u github.com/mikamai/karavel/cli`

### Secrets

Karavel delegates the handling of secrets and credentials to an external service. This approach has been chosen in order
to avoid having secrets stored in plain text inside Git repository or needing complex encryption mechanisms to protect them.

Karavel leverages the [External Secrets controller] to fetch credentials from the backing service and convert them into regular
Kubernetes `Secrets` objects to be consumed by pods. This way no changes to the application or special tools are required.

Credentials needed by core Karavel components are provisioned using this same method. This includes Git provider credentials for ArgoCD,
DNS providers tokens for managing DNS records and OpenID Connect client secrets for authentication.

The exact secrets required by each component are referenced in their specific page in the [Components section](./components/index.md).
They need to be provisioned beforehand and External Secrets need the appropriate permissions to access them, otherwise the platform will not be able
to run. They can be managed using whatever method is preferred, being it manually or via Infrastructure as Code tools like Terraform or CDK.

## Repository setup

Karavel builds on the principles of GitOps, whereby the entire state of the system is written
to git repositories and periodically synced with the live cluster.  
Let's prepare a new repository.

```bash
$ mkdir my-karavel-infra && cd my-karavel-infra
$ git init
Initialized empty Git repository in /home/user/my-karavel-infra/.git/
```

Now we are ready to generate the base Karavel manifests for deployment.

## Bootstrap

Each Karavel component is packaged and distributed as a [Helm chart]. The Karavel CLI tool
then uses these charts to generate the final YAML files that will be installed on the cluster. Ideally, as an end user
you will never need to interact with the charts directly.

Here is a super quick demo of the bootstrap process.

<script id="asciicast-389527" src="https://asciinema.org/a/389527.js" async></script>

The CLI tool uses a simple [HCL] file to describe the required components and their configuration. Components will also automatically
enable or disable  some optional features based on what other components are available. For example, a component may enable metrics collection if 
`prometheus` is present as well.

A component definition consists of a named block featuring the required version, target namespace and configuration params
that will be used to render out the manifests.

```hcl
component "name" {
  version = "major.minor.patch"
  namespace = "target-namespace"
  
  # Params
  
  param1 = true
  
  param2 = {
    example = "value"
  }
}
```

Here is an example with the [cert-manager] component:

```hcl
component "cert-manager" {
  version = "0.1.0"
  namespace = "cert-manager"

  letsencrypt = {
    email = "tech@karavel.io"
  }
}
```

Write your desired selection of components to a file called `karavel.hcl`. The `karavel render` command will use this file to assemble
a plan and render the proper configuration.

!!! info
    Instead of writing the `karavel.hcl` file from scratch, it is recommended to run `karavel init` to fetch the base config from the official repository.  
    The Karavel team regularly publishes recommended selections of components with matching versions that can be used as a starting point for your clusters.  
    To download a specific version of the Karavel Platform instead of the latest one, you can pass the `--version` flag to `karavel init` with the desired version.  
    See the [CLI reference](cli.md#karavel-init) for more information.

When you're satisfied with your configuration, simply running `karavel render` in the same directory as the file will download the required components from the repository
and generate the appropriate directory structure.

```bash
$ karavel render
Rendering new Karavel project with config file /home/user/my-karavel-infra/karavel.hcl

Rendering component 'goldpinger' 0.1.0 at my-karavel-infra/vendor/goldpinger
Rendering component 'grafana' 0.1.0 at my-karavel-infra/vendor/grafana
Rendering component 'tempo' 0.1.0 at my-karavel-infra/vendor/tempo
Rendering component 'calico' 0.1.0 at my-karavel-infra/vendor/calico
Rendering component 'external-dns' 0.1.0 at my-karavel-infra/vendor/external-dns
Rendering component 'argocd' 0.1.0 at my-karavel-infra/vendor/argocd
Rendering component 'prometheus' 0.1.0 at my-karavel-infra/vendor/prometheus
Rendering component 'ingress-nginx' 0.1.0 at my-karavel-infra/vendor/ingress-nginx
Rendering component 'loki' 0.1.0 at my-karavel-infra/vendor/loki
Rendering component 'cert-manager' 0.1.0 at my-karavel-infra/vendor/cert-manager
Rendering component 'external-secrets' 0.1.0 at my-karavel-infra/vendor/external-secrets
```

Here's a view of the generated folders.

```
.
├── applications
│  ├── argocd.yml
│  ├── aws-node-termination-handler.yml
│  ├── calico.yml
│  ├── cert-manager.yml
│  ├── external-dns.yml
│  ├── external-secrets.yml
│  ├── goldpinger.yml
│  ├── grafana.yml
│  ├── ingress-nginx.yml
│  ├── kustomization.yml
│  ├── loki.yml
│  ├── prometheus.yml
│  └── tempo.yml
├── karavel.hcl
├── kustomization.yml
├── projects
│  ├── infrastructure.yml
│  └── kustomization.yml
└── vendor
   ├── argocd
   ├── aws-node-termination-handler
   ├── calico
   ├── cert-manager
   ├── external-dns
   ├── external-secrets
   ├── goldpinger
   ├── grafana
   ├── ingress-nginx
   ├── loki
   ├── prometheus
   └── tempo

```

The `vendor` folder contains Kubernetes manifests for all Karavel components.
Files in this folder are not supposed to be edited by hand, as any change would be overwritten by future Karavel updates.
If you need to customize them follow the [Customization guide].

The `applications` directory contains ArgoCD application manifests that are used to register the vendored components with the GitOps engine.
These files can be safely edited to customize how ArgoCD manages them.

The `projects` directory contains ArgoCD project manifests. It will only contain one project, `infrastructure`, which is used
to group all installed Karavel applications. Feel free to add your own.

Finally, `kustomization.yml` is a Kustomize stack that references all the vendored components that are needed to bootstrap the cluster, 
as well as the application manifests, and can be used to install the entire platform in one go. Once the bootstrap components are up and running, ArgoCD will pick
the rest of the applications up and take care of deploying the rest of the stack.

To bootstrap the cluster, run the following command while connected to it:

```bash
kustomize build . | kubectl apply -f -
``` 

!!! warning
    Due to the way `kubectl` applies multiple manifests at once, there may be some race conditions between created resources.
    If you get some errors along the lines of `no matches for kind "AppProject" in version "argoproj.io/v1alpha1"` just rerun the command again.  
    Also notice that some custom resources provided by non-bootstrap components (e.g. `ServiceMonitor` provided by `prometheus`)
    will fail to install because their component is not part of the bootstrap process. This is fine and can be ignored. Once ArgoCD
    is up and running it will take care of updating the missing parts.

You can check that all bootstrap components are correctly deployed by running `kubectl get pods --all-namespaces`.

## Git Push

Now that the repository is ready we can commit and push it to the upstream remote.

```bash
git add --all . && git commit -m "Bootstrap new cluster"
git push origin master
```

Once pushed ArgoCD will automatically pick it up and start syncing the manifests with the cluster state.
New changes can be deployed by simply committing them and pushing to `origin`.

Congratulations, you are now running a full fledged Karavel instance!

## Next steps

### Updating components

Karavel publishes regular updates to its components, either to fix or improve their manifests or to introduce new features
and integrations between them. Updates are also released when a new Kubernetes version is published, to keep components compatible 
(see the [Karavel compatibility matrix]).  
You should really strive to maintain your clusters up to date with the latest Kubernetes release, and Karavel should be
updated consequently.  
To do so, simply change the required component versions and update the parameters if necessary based on
the updated documentation, then rerun the `karavel render` command. This command is idempotent, so it is safe to run multiple times.

[karavel]: https://github.com/mikamai/karavel/tree/master/cli
[git]: https://git-scm.com/
[cert-manager]: https://cert-manager.io
[kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
[helm]: https://helm.sh
[kustomize]: https://kubernetes-sigs.github.io/kustomize/installation/
[Python]: https://www.python.org/
[External Secrets controller]: https://github.com/external-secrets/kubernetes-external-secrets
[Helm chart]: https://helm.sh
[NGINX Ingress Controller]: https://kubernetes.github.io/ingress-nginx/
[ArgoCD Application]: https://argoproj.github.io/argo-cd/core_concepts/
[Customization guide]: operator-guides/customizing.md
[Kubernetes release cycle]: https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md
[Karavel compatibility matrix]: ./versioning.md#kubernetes-compatibility-matrix
[HCL]: https://www.terraform.io/docs/language/syntax/configuration.html
