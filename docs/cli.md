# Karavel CLI

The Karavel CLI tool is a utility that helps to manage GitOps repositories for installing Karavel on Kubernetes
clusters.

The tool is written in Go and packaged as a single static binary that runs on Linux.

## karavel init

```
Initialize a new Karavel project

Usage:
  karavel init [WORKDIR] [flags]

Flags:
      --checksum-url string   Override the official URL pointing to the Karavel config file checksum to download. Requires setting --config-url too
      --config-url string     Override the official URL pointing to the Karavel config file to download. Requires setting --checksum-url too
      --force                 Overwrite the config file even if it already exists
  -h, --help                  help for init
  -o, --output-file string    Karavel config file name to create (default "karavel.hcl")
  -v, --version string        Karavel Container Platform version to initialize (default "latest")

Global Flags:
      --colors   Enable colored logs (default true)
      --debug    Output debug logs
  -q, --quiet    Suppress all logs except errors
```

The Karavel team regularly publishes recommended selections of components with matching versions that can be used as a
starting point for your clusters.

These starting configurations are completely optional and can be tweaked to accomodate your specific setup, or you could
write your own from scratch. However, they are a useful starting point as the provided component versions combinations
have been carefully tested to ensure they are fully compatible with each other.

This command will download the recommended base config for the latest version of Karavel to the current directory:

```bash
$ cd /tmp/karavel 
$ karavel init
Initializing new Karavel latest project at /tmp/karavel

Fetching bootstrap config from https://github.com/projectkaravel/platform/releases/latest/download/karavel.hcl with checksum https://github.com/projectkaravel/platform/releases/latest/download/karavel.hcl.sha256

Downloading file karavel.hcl from https://github.com/projectkaravel/platform/releases/latest/download/karavel.hcl
Download completed in 1.995229ms

Downloading file karavel.hcl.sha256 from https://github.com/projectkaravel/platform/releases/latest/download/karavel.hcl.sha256
Download completed in 991.628µs

Checksum successfully validated. Writing config file to /tmp/karavel/karavel.hcl
```

To download a specific version of the Karavel Container Platform instead of the latest one, you can pass the `--version`
flag to `karavel init` with the desired version.

```bash
$ cd /tmp/karavel 
$ karavel init --version 0.1.0
Initializing new Karavel v0.1.0 project at /tmp/karavel

Fetching bootstrap config from https://github.com/projectkaravel/platform/releases/v0.1.0/download/karavel.hcl with checksum https://github.com/projectkaravel/platform/releases/v0.1.0/download/karavel.hcl.sha256

Downloading file karavel.hcl from https://github.com/projectkaravel/platform/releases/v0.1.0/download/karavel.hcl
Download completed in 1.995229ms

Downloading file karavel.hcl.sha256 from https://github.com/projectkaravel/platform/releases/v0.1.0/download/karavel.hcl.sha256
Download completed in 991.628µs

Checksum successfully validated. Writing config file to /tmp/karavel/karavel.hcl
```

You can now safely edit the `karavel.hcl` file to configure the platform based on your environment. There are a few
parameters that some component need in order to properly function, like OIDC and Cloud providers credentials. More
information is provided in the next pages.

## karavel render

```

Render a Karavel project with the given config (defaults to 'karavel.hcl' in the current directory).

This command is idempotent and can be run multiple times without issues. 
It will respect changes made to files outside the 'vendor' directory, only adding or removing Karavel-specific entries.
It will, however, consider the 'vendor' directory as a fully-managed folder and may add, delete or modify any file inside it without warning.

Usage:
  karavel render [flags]

Flags:
  -f, --file string   Specify an alternate config file (default "karavel.hcl")
  -h, --help          help for render

Global Flags:
      --colors   Enable colored logs (default true)
      --debug    Output debug logs
  -q, --quiet    Suppress all logs except errors
```

`karavel render` is the primary tool used to manage Karavel GitOps repositories. It takes a [HCL] configuration file
that describes the required components, their version and their parameters, and generates the appropriate Kubernetes
manifests for them.  
These manifests will install all the required components in one go from a set of curated packages maintained by the
Karavel project.

The command will also enable or disable bits of configuration based on the available components, enabling useful
integrations without the user having to worry about wiring all the pieces together on its own.  
One common example would be adding [ServiceMonitor] definitions to all the components that provide Prometheus metrics if
the `prometheus` component is added to the configuration, so that metrics can be scraped and visualized in Grafana.

Once bootstrapped a new cluster is completely self-managed, with [ArgoCD] acting as the GitOps engine in charge of
maintaining the deployed services in sync with their manifests stored in your Git provider of choice.

Changing a configuration and redeploying becomes as easy as committing and pushing to a Git repository.

Check the [Quickstart Guide] to see this tool in action.

[ArgoCD]: https://argoproj.github.io/argo-cd

[Quickstart Guide]: quickstart.md

[HCL]: https://www.terraform.io/docs/language/syntax/configuration.html

[ServiceMonitor]: https://github.com/prometheus-operator/prometheus-operator/blob/master/Documentation/user-guides/getting-started.md
