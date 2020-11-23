# Example Bootstrap

This folder contains an example infrastructure stack
that is capable of bootstrapping a fully configured Karavel Platform instance.

It uses the `mikamai.karavel` Ansible collection to render all the necessary Kubernetes manifests
and perform some preliminary setup, like creating namespaces.

Two playbooks are provided:

- [bootstrap.yml](./bootstrap.yml) should be run the first time a new repo is set up.  
  It will upsert the required namespaces and generate the platform manifests.
- [update.yml](./update.yml) can be run afterwards to update the platform manifests, for example when upgrading to a new
  version of Karavel. It will still upsert the required namespaces, but this step can be skipped by setting the `init_namespaces` variable to `false`.

The bootstrap tool will generate the following files and directories:

- `applications/` contains [ArgoCD applications](https://argoproj.github.io/argo-cd/core_concepts) that manage platform components.  
  You are free to add your own by simply adding a new `Application` file and referencing it in `applications/kustomization.yml`.
- `vendor/` contains [Kustomize](https://kustomize.io) stacks that install the various Karavel components.  
  You are not supposed to edit these files by hand as they will be overwritten when running the `update.yml` playbook. 
  Instead you should create [Kustomize overlays](https://kubectl.docs.kubernetes.io/references/kustomize/glossary/#overlay) that patch the provided resources with your changes.
- `kustomization.yml` is a one-off Kustomize stack that can be used to bootstrap the platform for the first time. It references all the stacks in the other folders to easily
  install the whole platform at once. It can be deleted afterwards since `update.yml` will not recreate it.

## Configuration

Karavel needs a few config parameters to setup specific features like
DNS records, secrets names and various credentials. These parameters are documented in 
`vars.all.yml`. You should copy this file to `vars.yml` and edit it with the correct information
for your setup.

## First boostrap

```bash
# Edit variables with the correct information for your setup
cp vars.all.yml vars.yml
$EDITOR vars.yml

# Run the bootstrap tool
ansible-playbook bootstrap.yml

# Bootstrap the platform
kustomize build . | kubectl apply -f -

# Cleanup (optional)
rm kustomization.yml bootstrap.yml

# Add everything to Git
git add --all . && git commit -m "Bootstrap Karavel"

# Push to origin
git push origin master

# ArgoCD will take over from now on. `kubectl apply` is no longer needed
```

## Update components

Before updating Karavel it is recommended to read the Changelog and Update Guides
to ensure a smooth upgrade process.

// TODO: reference update guides and changelogs when we have them

```bash
# Update components
ansible-playbook update.yml

# Add everything to Git
git add --all . && git commit -m "Update Karavel"

# Push to origin
git push origin master

# ArgoCD will sync the manifests and apply the updates
```
