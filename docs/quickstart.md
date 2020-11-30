# Quickstart

This section will guide you through the first bootstrap
of a new Karavel cluster from scratch. It assumes you have a general knowledge
of Git and Kubernetes and know your way around a command shell.

## Requirements

### Tools

- [git], [ansible], [kubectl], [helm] and [kustomize] installed on your local machine
- admin access to a running Kubernetes cluster. Minikube is fine for a simple, local deployment for testing purposes.
- Ansible will also need [Python] to run, as well as the `openshift` module to interact with the cluster.
  It can be installed by running `pip install openshift`.

### Secrets

Karavel delegates the handling of secrets and credentials to an external service. This approach has been chosen in order
to avoid having secrets stored in plain text inside Git repository or needing complex encryption mechanisms to protect them.

Karavel leverages the [External Secrets controller] to fetch credentials from the backing service and convert them into regular
Kubernetes `Secrets` objects to be consumed by pods. This way no changes to the application or special tools are required.

Credentials needed by core Karavel components are provisioned using this same method. This includes Git provider credentials for ArgoCD,
DNS providers tokens for managing DNS records and OpenID Connect client secrets for authentication.

The exact secrets required by each component are referenced in the [Bootstrap variables reference] section of the Operator Handbook.
They need to be provisioned beforehand and External Secrets need the appropriate permissions to access them, otherwise the platform will not be able
to run.

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

Each Karavel component is packaged and distributed as a [Helm chart]. The Karavel bootstrap tool
then uses these charts to generate the final YAML files that will be installed on the cluster. Ideally, as an end user
you will never need to interact with the charts directly.

The bootstrap tool is written using [ansible] and is published on [Ansible Galaxy](https://galaxy.ansible.com/mikamai/karavel) as a [Collection](https://docs.ansible.com/ansible/latest/user_guide/collections_using.html) of roles
that will setup the GitOps repository for you.

Let's define an Ansible playbook to run the first cluster bootstrap. This step will only be needed the first time a cluster is set up.

Add a file called `bootstrap.yml` to the root of the repository with the following content.

```yaml
- hosts: localhost
  vars_files:
    - vars.yml
  tasks:
    - import_role:
        name: mikamai.karavel.core
    - import_role:
        name: mikamai.karavel.nginx
    - import_role:
        name: mikamai.karavel.argo_apps
    - import_role:
        name: mikamai.karavel.bootstrap_kustomization
```

Each role imported from the `mikamai.karavel` collection prepares a set of components for deployment.  
`mikamai.karavel.core` installs the core components that are fundamental for a Karavel cluster to operate properly.  
`mikamai.karavel.nginx` installs the [NGINX Ingress Controller] to provide routing and load balancing of application through the Public Internet.  
`mikamai.karavel.argo_apps` prepares the [ArgoCD Application] definitions that will register Karavel with the GitOps engine.  
`mikamai.karavel.bootstrap_kustomization` generates a `kustomization.yml` that can be used to manually install the entire platform on a fresh cluster in a single command.

The bootstrap playbook references a `var.yml` files. This file contains the configuration parameters that you will need to fill out with information that are specific to your environment.
It contains variables for DNS domains to use, credential secrets to pull from the secure vault and provider configuration for some of the interchangeable services the platform supports.
See the [Bootstrap variables reference] for information on all the available variables.

Once all the appropriate variables has been filled out it is time to bootstrap Karavel. Running the tool is as simple as running a regular Ansible playbook.

```bash
ansible-playbook bootstrap.yml
```

This will generate a few directories and files in the repository. These directories will always be generated relative to the playbook file, so if you wish to keep them in a subdirectory
you can simply move the `bootstrap.yml` file in it and run it from there.

Here's a view of the generated folders.

```
.
|-- applications
|   |-- argocd.yml
|   |-- bootstrap.yml
|   |-- calico.yml
|   |-- cert-manager.yml
|   |-- dex.yml
|   |-- external-dns.yml
|   |-- external-secrets.yml
|   |-- goldpinger.yml
|   |-- ingress-nginx.yml
|   |-- kustomization.yml
|   `-- project.yml
|-- bootstrap.yml
|-- kustomization.yml
|-- vars.yml
`-- vendor
    |-- argocd
    |-- calico
    |-- cert-manager
    |-- dex
    |-- external-dns
    |-- external-secrets
    |-- goldpinger
    `-- ingress-nginx
```

The `vendor` folder contains Kubernetes manifests for all Karavel components.
Files in this folder are not supposed to be edited by hand, as any change would be overwritten by future Karavel updates.
If you need to customize them follow the [Customization guide].

The `applications` directory contains ArgoCD application manifests that are used to register the vendored components with the GitOps engine.
These files can be safely edited to customize how ArgoCD manages them.

Finally, `kustomization.yml` is a Kustomize stack that references all the vendored components as well as the application manifests, and can be used to install
the entire platform in one go.
To do so, run the following command while connected to the cluster:

```bash
kustomize build . | kubectl apply -f -
``` 

!!! warning
    At the time of writing (2020-11-26) there may be some race conditions between created resources.
    If you get some errors along the lines of `no matches for kind "AppProject" in version "argoproj.io/v1alpha1"`
    just rerun the same command again.

You can check that all services are correctly deployed by running `kubectl get pods --all-namespaces`.

## Prepare for updates

Karavel publishes regular updates that follow the [Kubernetes release cycle].
You should really strive to maintain your clusters up to date with the latest Kubernetes release, and Karavel should be
updated consequently. To do so we can write a second playbook that only updates the vendored module, skipping the bootstrap parts (which we assume have already
been performed).

Add a file called `update.yml` to the root of the repository with the following content.

```yaml
- hosts: localhost
  vars_files:
    - vars.yml
  tasks:
    - import_role:
        name: mikamai.karavel.core
    - import_role:
        name: mikamai.karavel.nginx
```

Additionally, you can set the `init_namespaces` parameter to `false` in `vars.yml` to avoid re-ensuring that 
the required namespaces exist on the cluster. You can keep it to `true` if you want, nothing bad is gonna happen as
the bootstrap tool will see that the namespaces are already present and ignore them, but it can speed up the update process a little bit.

To update the vendored components simply run `ansible-playbook update.yml`. No need to manually apply them with `kubectl` though: they will be automatically
synced by ArgoCD.
 
## Git Push

Now that the repository is ready we can commit and push it to the upstream remote.

```bash
git add --all . && git commit -m "Bootstrap new cluster"
git push origin master
```

Once pushed ArgoCD will automatically pick it up and start syncing the manifests with the cluster state.
New changes can be deployed by simply committing them and pushing to `origin`.

Congratulations, you are now running a full fledged Karavel instance!

[git]: https://git-scm.com/
[ansible]: https://ansible.com/
[kubectl]: https://kubernetes.io/docs/tasks/tools/install-kubectl/
[helm]: https://helm.sh
[kustomize]: https://kubernetes-sigs.github.io/kustomize/installation/
[Python]: https://www.python.org/
[External Secrets controller]: https://github.com/external-secrets/kubernetes-external-secrets
[Helm chart]: https://helm.sh
[NGINX Ingress Controller]: https://kubernetes.github.io/ingress-nginx/
[ArgoCD Application]: https://argoproj.github.io/argo-cd/core_concepts/
[Bootstrap variables reference]: /operator-handbook/variables
[Customization guide]: /operator-handbook/customizing
[Kubernetes release cycle]: https://github.com/kubernetes/sig-release/blob/master/releases/release_phases.md
