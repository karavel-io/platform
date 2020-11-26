# Bootstrap tool

The bootstrap tool consists of a collection of [Ansible] roles that 
take care of generating the necessary Kubernetes manifests to install Karavel on a new cluster.

It can be found in the [platform/bootstrap] directory of the main GitHub repository.

## Requirements

Ansible 2.10 or newer is required. 
The [ansible-galaxy] tool is included with every Ansible installation.
 
## Local Install

The roles are packaged as an [Ansible Collection] that can be bundled up and published on [Ansible Galaxy].
To test it locally however, the following commands can be run to build and install it.

```bash
$ ansible-galaxy collection build -f
$ ansible-galaxy collection install mikamai-karavel-$VERSION.tar.gz -f
```

The collection's roles can now be referenced by any playbook.

[Ansible]: https://ansible.com
[platform/bootstrap]: https://github.com/mikamai/karavel/tree/master/platform/bootstrap
[ansible-galaxy]: https://docs.ansible.com/ansible/latest/cli/ansible-galaxy.html
[Ansible Collection]: https://docs.ansible.com/ansible/latest/user_guide/collections_using.html
[Ansible Galaxy]: https://galaxy.ansible.com
