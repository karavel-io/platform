# Karavel Bootstrap Tool

The Karavel Bootstrap Tool is a utility that helps scaffolding
new GitOps repositories for installing Karavel on a Kubernetes cluster.

The tool is written with [Ansible] and packaged as an [Ansible Collection].
This collection contains Ansible roles that will generated Kubernetes manifests in a Git repository.

These manifests will install all the required components in one go from a set of curated packages maintained
by the Karavel project. Once bootstrapped a new cluster is completely self-managed, with [ArgoCD] acting as the GitOps engine
in charge of maintaing the deployed services in sync with their manifests stored in your Git provider of choice.

Changing a configuration and redeploying becomes as easy as committing and pushing to a Git repository.

Check the [Quickstart Guide] to see this tool in action.

[Ansible]: https://ansible.com
[Ansible Collection]: https://docs.ansible.com/ansible/latest/user_guide/collections_using.html
[ArgoCD]: https://argoproj.github.io/argo-cd
[Quickstart Guide]: ../quickstart.md
