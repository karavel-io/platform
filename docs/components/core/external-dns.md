# External DNS

[![Component version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=component+version&query=$.entries['external-dns'][0].version&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](./external-dns.md)
[![External DNS version](https://img.shields.io/badge/dynamic/yaml?color=blue&label=external-dns+version&query=$.entries['external-dns'][0].appVersion&url=https%3A%2F%2Fcharts.mikamai.com%2Fkaravel%2Findex.yaml&style=for-the-badge)](https://github.com/kubernetes-sigs/external-dns)
![Bootstrap component](https://img.shields.io/badge/bootstrap-true-orange?style=for-the-badge)

## Overview

ExternalDNS synchronizes exposed Kubernetes Services and Ingresses with DNS providers.  
In a dynamic and elastic environment like Kubernetes, services come and go in the blink of an eye. Developers 
may want deploy a new service to the public and need a DNS record to route traffic into the cluster. Having to
ask an administrator to manually configure the DNS provider can be tedious and a waste of time, especially in
large organizations. External DNS bridges the gap between Kubernetes network objects, like `Service` and `Ingress` definitions,
and configure the upstream DNS provider accordingly, taking care of creating, updating or deleting records as needed.

Multiple DNS providers are supported. At the moment the following services can be configured:

- [AWS Route53]
- [Cloudflare]

Support for additional services can be added based on the controller's capabilities.

## Component configuration

```hcl
component "external-dns" {
  version = "0.1.0"
  namespace = "external-dns"

  # Params default values

  # when configured with a base domain, external-dns will ignore requests that are not children domains
  domainFilter = ""
  
  # Upstream DNS provider to configure
  # required, must be one of 'cloudflare', 'route53'
  provider = "" 
  
  cloudflare = {
    # Enable or disable the Cloudflare Proxy on managed records. Can be overridden on a per-object basis
    proxied = false
    
    # Restrict to domains in a specific Cloudflare Zone. Optional
    zoneId = ""
    
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
  
  route53 = {
    # Only look at zone of this type (values can be 'public', 'private' or empty for both)
    zoneType = ""
    # Restrict to domains in a specific Route53 Zone. Optional
    zoneId = ""
    # Configure when deployed on EKS or other platforms with IAM Roles for Service Accounts
    eksRole = ""
    # Configure when deployed on AWS with KIAM
    iamRole = ""
  }
}
```

[AWS Route53]: https://aws.amazon.com/route53/
[Cloudflare]: https://cloudflare.com
