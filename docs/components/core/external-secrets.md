# External Secrets

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries['external-secrets'][0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./external-secrets.md)
[![External Secrets version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=external-secrets+version&query=$.entries['external-secrets'][0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://github.com/kubernetes-sigs/external-secrets)
![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge)

## Overview

The most critical aspect of every kind of infrastructure is secrets management. Secure storage and handling of
passwords, private keys and other sensitive credentials is of primary importance for every production-grade system
and should be treated with outmost carefulness.

The Karavel team took the decision of delegating secrets management to an external vault right from the start.
Instead of keeping credentials in plain text in the Git repository, or using complicated encryption mechanism to 
keep them from being abused, Karavel connects to an external service to securely retrieve them and bind them to regular
[Kubernetes Secret] objects so that they can be injected into pod's environment variables or mounted as files.  
This means that no changes have to be made to applications that already consume their secrets via envars or config files.

The Karavel External Secrets component uses the [Kubernetes External Secrets] controller to poll the external vault
and synchronize information between the store and Kubernetes.

Multiple secure vaults are supported. At the moment the following services can be configured:

- [AWS Secrets Manager]
- [Hashicorp Vault]

Support for additional services can be added based on the controller's capabilities.

## Component configuration

```hcl
component "external-secrets" {
  version = "version"
  namespace = "namespace"
  
  # Params default values

  pollingIntervalMs = 300000

  aws = {
    enable = false
    region = ""
    defaultRegion = "eu-west-1"

    # Configure when deployed on EKS or other platforms with IAM Roles for Service Accounts
    eksRole = ""
    # Configure when deployed on AWS with KIAM
    iamRole = "" 
  }

  vault = {
    enable = false
    address = ""
    defaultMountPoint = ""
    defaultRole = ""
    extraCertsSecret = {
      name = "vault-ca"
      key = "ca.pem"
    }
  }
}
```

[Kubernetes Secret]: https://kubernetes.io/docs/concepts/configuration/secret/
[Kubernetes External Secrets]: https://github.com/external-secrets/kubernetes-external-secrets
[AWS Secrets Manager]: https://aws.amazon.com/secrets-manager/
[Hashicorp Vault]: https://vaultproject.io/
