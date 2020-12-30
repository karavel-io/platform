# Bootstrap variables reference

The [bootstrap tool] requires a few configuration variables to inject environment-dependant
information into services. This includes various credentials for external systems and other 
values that are specific to the organization that is running Karavel.

This section documents all the available variables for each component. All variables
should be written to the `vars.yml` file that is consumed by the bootstrap tool.

## Bootstrap

The following variables are used to customize the bootstrap tool itself.

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `init_namespaces` | boolean | `true` | If `true`, the bootstrap tool will initialize namespaces on the live cluster. This can be turned off to only set the local repository up. |
| `application_dir` | path | `applications` | Where the ArgoCD applications manifests will be written. |
| `vendor_dir` | path | `vendor` | Where the Karavel components manifests will be written. |

### Example

```yaml
# vars.yml

init_namespaces: true 
application_dir: applications
vendor_dir: vendor
```

## Calico

[Calico] is the network plugin used by Karavel. 
The following variables are located under the `calico` section (see example).

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| enable_cni | boolean | true | If set to `false` Calico will not install the CNI plugin to manage networking and will only work for [Network Policies]. |
| ha_setup | boolean | false | If set to `true` Calico will be deployed in [High Availability mode](https://docs.projectcalico.org/reference/typha/overview) with Typha

### Example

```yaml
# vars.yml

calico:
    enable_cni: true
    ha_setup: false
```

## External Secrets

[External Secrets] is the controller that manages the integration between the platform and the backing secret vault. Currently the supported backends are:

- [AWS Secrets Manager] (with `provider: aws`)
- [AWS Systems Manager Parameter Store] (with `provider: aws`)
- [Hashicorp Vault] (with `provider: vault`)

### General configuration

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `polling_interval_ms` | number | `300000` (5 minutes) | On some pay-as-you-go services like AWS SM lowering this value can bring higher costs due to the increased API traffic to the service. |
| `provider` | string | none (required) | Must match one of the available providers |

#### Example

```yaml
# vars.yml

secrets:
    polling_interval_ms: 300000
```

### AWS Secrets Manager and AWS Systems Manager Parameter Store

Used in combination `provider: aws`.

These backends share the same configuration. They are used in secrets with `backend: secretsManager` and `backend: systemManager` respectively.

Follow the [External Secrets] documentation for more information regarding AWS IAM permissions required for the IAM roles.

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `region` | string | none (required) | Must match the [AWS Region] of the target secrets.
| `default_region` | string | none (required) | Must match the [AWS Region] of the target secrets.
| `eks_role` | string | none (required) | AWS IAM Role that will be used by the External Secrets controller to authenticate with AWS. Only used when running on [EKS].
| `iam_role` | string | none (required) | AWS IAM Role that will be used by the External Secrets controller to authenticate with AWS. Only used when running on EC2 with the kiam setup.

#### Example

```yaml
# vars.yml

secrets:
      # other config params
    
      provider: aws
      region: eu-west-1
      default_region: eu-west-1
      eks_role: arn:aws:iam::1234567890:role/KaravelExampleSecrets  # When running on EKS
      iam_role: arn:aws:iam::1234567890:role/KaravelExampleSecrets  # When running on EC2
```

### Hashicorp Vault

Used in combination `provider: vault`.

This provider uses the [Vault Kubernetes Auth Method] to access the Vault cluster.

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `address` | url | none (required) | URL of the Vault API | 
| `default_mount_point` | string | none (optional) | Default secret mount point if not specified in the `ExternalSecret` objects |
| `default_role` | string | none (optional) | Default Vault role to use if not specified in the `ExternalSecret` objects |
| `extra_certs_secret_ref` | SecretRef | none (optional) | Reference to a Kubernetes secret containing the Vault custom CA certificate. Only required if Vault is using a custom CA. 

#### Example

```yaml
# vars.yml

secrets:
    # other config params

    provider: vault
    address: https://vault.example.com:8200
    default_mount_point: my-karavel-cluster
    default_role: my-k8s-role
    extra_certs_secret_ref:
      name: vault-ca
      key: ca.pem
```

## External DNS

Karavel includes the [ExternalDNS] controller for automatically provisioning DNS records pointing to cluster resources.
It currently supports the following DNS providers:

- [CloudFlare] (with `provider: cloudlfare`)
- [Amazon Route53] (with `provider: route53`)

### General configuration

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `domain_filter` | hostname | none (optional) | Limit ExternalDNS to only operate on hosted zones that match this domain. Omit to process all hosted zones |
| `provider` | string | none (required) | Must match one of the available providers |

#### Example

```yaml
# vars.yml

dns:
  domain_filter: example.com
```

### CloudFlare

Configure ExternalDNS to create records on CloudFlare. 
It requires generating a [CloudFlare API Token] that must be stored on the [secure vault](#external-secrets).

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `proxied` | boolean | `true` | If the newly created records should have [CloudFlare Proxying] enabled or not. Can be overridden on a [per-ingress basis](https://github.com/kubernetes-sigs/external-dns/blob/master/docs/tutorials/cloudflare.md#setting-cloudflare-proxied-on-a-per-ingress-basis) |
| `hosted_zone_id` | string | none (optional) | Limit ExternalDNS to a specific CloudFlare hosted zone |
| `secret_ref` | ExternalSecretRef | none (required) | Configuration for the `ExternalSecret` object that will fetch the CloudFlare API token. |

#### Example

```yaml
# vars.yml

dns:
  # other config params

  provider: cloudflare
  proxied: true
  hosted_zone_id: ABCDEXAMPLE
  secret_ref:
    backend: secretsManager             # one of the supported backends
    key: my-karavel-cluster/cloudflare  # key for the external secret on the backend
    property: api-token                 # property inside the external secret that contains the CloudFlare API token
```

### Amazon Route53

Configure ExternalDNS to create records on Amazon Route53. 
It requires creating an AWS IAM Role that can access Route53. 

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `region` | string | none (required) | Must match the [AWS Region] of the target hosted zones.
| `hosted_zone_id` | string | none (optional) | Limit ExternalDNS to a specific Route53 hosted zone |
| `hosted_zone_type` | string | none (optional) | Limit ExternalDNS to a specific Route53 hosted zone type. Can be `public` or `private`. Leave empty for `both`. |
| `eks_role` | string | none (required) | AWS IAM Role that will be used by the ExternalDNS controller to authenticate with AWS. Only used when running on [EKS].
| `iam_role` | string | none (required) | AWS IAM Role that will be used by the ExternalDNS controller to authenticate with AWS. Only used when running on EC2 with the kiam setup.

#### Example

```yaml
# vars.yml

dns:
    # other config params

    provider: route53
    region: eu-west-1
    hosted_zone_id: ABCDEXAMPLE
    hosted_zone_type: ""
      eks_role: arn:aws:iam::1234567890:role/KaravelExampleSecrets  # When running on EKS
      iam_role: arn:aws:iam::1234567890:role/KaravelExampleSecrets  # When running on EC2
```

## cert-manager

[cert-manager] is the cluster component used to manage TLS certificates. It can create a wide variety of different certificates, both
self-signed or with a private CA, or leverage an external [ACME] provider such as [Let's Encrypt].

Karavel configures a Let's Encrypt issuer by default so that valid TLS certificates can be generated to expose
public HTTPS endpoints.

cert-manager uses the [http01] ACME challenge by default, but can also leverage the [dns01] challenge when integrated with
a DNS provider.

It currently supports the following DNS providers:

- [CloudFlare] (with `provider: cloudlfare`)

### General configuration

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `email` | email | none (required) | Email that will be associated with the Let's Encrypt account for certificate issuing |
| `provider` | string | none (required) | Must match one of the available providers |

#### Example

```yaml
# vars.yml

letsencrypt:
  email: tech@example.com
```

### CloudFlare

When configured as the DNS provider, CloudFlare records will be created for each ACME challenge initiated by cert-manager.
If CloudFlare has been configured as the [ExternalDNS provider](#external-dns) too, this configuration section can be skipped and the 
same parameters will be used for both components.

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `secret_ref` | ExternalSecretRef | none (required) | Configuration for the `ExternalSecret` object that will fetch the CloudFlare API token. |

#### Example

```yaml
# vars.yml

letsencrypt:
  # other config params

  provider: cloudflare
  secret_ref:
    backend: secretsManager             # one of the supported backends
    key: my-karavel-cluster/cloudflare  # key for the external secret on the backend
    property: api-token                 # property inside the external secret that contains the CloudFlare API token
```

## Dex

TODO

## ArgoCD

[ArgoCD] is the beating heart of Karavel, the GitOps engine monitoring and deploying applications
and services to the cluster.

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `public_host` | hostname | none (required) | Public host for ArgoCD UI and API |
| `git.repo` | Git HTTPS/SSH URL | none (required) | Git URL for the repository holding Karavel manifests |
| `oidc` | object | none (required) | [ArgoCD OIDC configuration object]. **NOTICE** This object is passed as-is, so keys should be camelCased instead of snake_cased |
| `credential_secret` | ExternalSecretRef | none (required) | Configuration for the `ExternalSecret` object that will fetch the Git provider credentials used to pull the Git repository |
| `secret` | ExternalSecretRef | none (required) | Configuration for the `ExternalSecret` object that will fetch the [ArgoCD secret object] |

### Example

```yaml
# vars.yml

argocd:
  public_host: argocd.example.com
  git:
    # Git repository that holds the platform manifests
    # A.K.A. this repo
    repo: git@github.com:example/example.git

  # OIDC provider for authentication
  oidc:
    config:
      name: "Sample"
      issuer: "https://sample.auth0.com"
      client_id: argocd
      requested_scopes: [ "openid", "profile", "email", "groups" ]

  # creates the infrastructure-repo-secret from the secure backend
  # Must contain the keys `username` and `password` for Git basic auth/token based auth
  # OR
  # Must contain the key `sshPrivateKey` for Git SSH auth
  credentials_secret:
    backend: secretsManager
    key: my-cluster/argocd-pull-creds
    type: ssh

  # creates the argocd-secret manifest from the secure backend
  # Can contain all the keys defined in argocd-secrets
  # https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/argocd-secret.yaml#L10
  secret:
    backend: secretsManager
    key: my-cluster/argocd-secret
```

## AWS Node Termination Handler

The [AWS Node Termination Handler] is a small daemon that manages rebooting EC2 instances used as cluster nodes.
It properly cordons and drains nodes before they are terminated or rebooted so that Kubernetes can cleanly reschedule them.

This component is only needed in clusters running on AWS EC2 or EKS.

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `processor` | string | imds | Chooses the underlying processor. At the moment only `imds` for the Instance Metadata Service processor is supported. |
| `webhook.url` | url | none (optional) | Webhook URL that will receive termination updates. |
| `webhook.headers` | object | `{ "Content-type": "application/json" }` | Headers sent to the Webhook URL |
| `webhook.tpl` | url | see below | Message template sent to the Webhook URL |

Default Webhook template

```
[NTH][Instance Interruption]
EventID: {{ .EventID }}  - Kind: {{ .Kind }} - Instance: {{ .InstanceID }} - Start Time: {{ .StartTime }}
{{ .Description }}
```

### Example

```yaml
# vars.yml

aws_node_termination_handler:
  processor: imds
  webhook:
    url: "https://example.com/slack"
    headers:
      Content-type: application/json
      Authorization: Bearer <token>
    tpl: |
      [NTH][Instance Interruption]
      EventID: {{ .EventID }}  - Kind: {{ .Kind }} - Instance: {{ .InstanceID }} - Start Time: {{ .StartTime }}
      {{ .Description }}
```

## Grafana

[Grafana] is the core of the observability platform offered by Karavel. It aggregates all the observability tools
for logging, metrics and tracing and present them through a unified web interface. Karavel deploys the [grafana-operator]
so that dashboards and datasources can be managed as Kubernetes resources and deployed alongside services and deployments.


| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `host` | hostname | none (required) | Public host for Grafana |
| `secret` | ExternalSecretRef | none (required) | Configuration for the `ExternalSecret` object that will fetch the [Grafana] environment variables. Keys will be passed as-is as envar names |

### Example

```yaml
grafana:
  host: grafana.example.com
  secret:
    backend: secretsManager
    key: my-cluster/grafana-secret
```

## Loki

[Loki] is a logging aggregation tool by Grafana that is cheap and fast to operate. Karavel allows to configure Loki
with a local filesystem, or an S3 bucket as the backing storage. It will configure itself as a Grafana datasource
so that it will be available in the cluster Grafana. It currently supports the following stores:

- Filesystem (with `store: filesystem`)
- Amazon S3 or compatible (with `store: s3`)

### General configuration

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `store` | string | none (required) | Must match one of the available stores |

#### Example

```yaml
loki:
  store: filesystem
```

### S3

Loki can store indexes and chunks in an S3 bucket or a compatible implementation, like Minio or Ceph.

| Variable | Type | Default | Notes |
| -------- | ---- | ------- | ----- |
| `bucket` | string | none (required) | S3 bucket to use as the backing storage |
| `endpoint` | hostname | none (optional) | Custom S3 endpoint when using a non-AWS S3 implementation (e.g. Minio) |
| `encrypted` | boolean | `true` | Use SSE Encryption |
| `insecure` | boolean | `false` | Support insecure (plain-text) connections to S3 (e.g. for local Minio) |
| `path_style` | boolean | `false` | Use path-style instead of virtual-host strategy for constructing bucket URLs (e.g. for Minio) |
| `region` | string | none (required) | Must match the [AWS Region] of the cluster to improve speed. |
| `eks_role` | string | none (required) | AWS IAM Role that will be used by Loki to authenticate with AWS. Only used when running on [EKS]. |
| `iam_role` | string | none (required) | AWS IAM Role that will be used by Loki to authenticate with AWS. Only used when running on EC2 with the kiam setup. |

### Example

```yaml
loki:
  store: s3
  s3:
    bucket: my-logging-bucket
    endpoint: my-minio.example.com
    region: eu-west-1
    encrypted: true
    insecure: false
    path_style: false
    eks_role: arn:aws:iam::1234567890:role/KaravelExampleSecrets  # When running on EKS
    iam_role: arn:aws:iam::1234567890:role/KaravelExampleSecrets  # When running on EC2
```

[bootstrap tool]: ./bootstrap.md
[Calico]: https://projectcalico.org/
[Network Policies]: https://kubernetes.io/docs/concepts/services-networking/network-policies/
[External Secrets]: https://github.com/external-secrets/kubernetes-external-secrets/
[AWS Secrets Manager]: https://aws.amazon.com/secrets-manager/
[AWS Systems Manager Parameter Store]: https://docs.aws.amazon.com/systems-manager/latest/userguide/systems-manager-parameter-store.html
[Hashicorp Vault]: https://vaultproject.io/
[AWS Region]: https://aws.amazon.com/about-aws/global-infrastructure/regions_az/
[EKS]: https://aws.amazon.com/eks/
[Vault Kubernetes Auth Method]: https://www.vaultproject.io/docs/auth/kubernetes/
[External DNS]: https://github.com/kubernetes-sigs/external-dns/
[CloudFlare]: https://cloudflare.com/
[Amazon Route53]: https://aws.amazon.com/route53/
[CloudFlare API Token]: https://developers.cloudflare.com/api/tokens/create
[CloudFlare Proxying]: https://support.cloudflare.com/hc/en-us/articles/205177068
[cert-manager]: https://cert-manager.io/
[ACME]: https://en.wikipedia.org/wiki/Automated_Certificate_Management_Environment
[Let's Encrypt]: https://letsencrypt.org/
[http01]: https://cert-manager.io/docs/configuration/acme/http01/
[dns01]: https://cert-manager.io/docs/configuration/acme/dns01/
[ArgoCD]: https://argoproj.github.io/argo-cd
[ArgoCD OIDC configuration object]: https://argoproj.github.io/argo-cd/operator-manual/user-management/#existing-oidc-provider
[ArgoCD Secret object]: https://github.com/argoproj/argo-cd/blob/master/docs/operator-manual/argocd-secret.yaml#L10
[AWS Node Termination Handler]: https://github.com/aws/aws-node-termination-handler
[Grafana]: ./components/grafana.md
[grafana-operator]: https://github.com/integr8ly/grafana-operator
[Loki]: ./components/logging.md
