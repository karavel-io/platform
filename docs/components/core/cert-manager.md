# cert-manager

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries['cert-manager'][0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./cert-manager.md)
[![cert-manager version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=cert-manager+version&query=$.entries['cert-manager'][0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://cert-manager.io)
![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge)

## Overview

Automates TLS certificates provisioning from various sources, including ACME (Let's Encrypt), HashiCorp Vault, Venafi, self signed and private certificate authorities.

[cert-manager] is a Kubernetes controller that manages TLS certificate on behalf of applications, automating the provisioning and renewal
of public and private certificates for encrypting connections.

## Component configuration

```hcl
component "$NAME" {
  version = "version"
  namespace = "namespace"

  # Params default values

  # Configure how Let's Encrypt certificates are provisioned
  letsencrypt = {
    # Email associated to the Let's Encrypt account 
    # that will be created to request certificates
    # It is used by LE to send notifications regarding 
    # certificate expiration if cert-manager doesn't renew them in time
    email = ""
    
    # Configuration for using Cloudflare for the dns01 challenge
    cloudflare = {
      enable = false
      
      # Email of the target Cloudflare account
      email = ""

      # ExternalSecret object reference to a secrets holding the Cloudflare API Token
      secret = {
        # Must be one of the services configured in the External Secrets component
        backend = "secretsManager"
        
        # Backend-specific key for the target secret
        key = ""
        
        # Optional nested property inside the upstream secret
        property = ""
      }
    }
  }
}
```

[cert-manager]: https://cert-manager.io
