# Karavel CLI

The Karavel CLI tool is a utility that helps to manage GitOps repositories for installing Karavel on Kubernetes
clusters.

## Install

Binaries for all mainstream operating systems can be downloaded from [GitHub](https://github.com/karavel-io/cli/releases).

### Nix

The CLI is packaged as a [Flake](https://nixos.wiki/wiki/Flakes). You can run it as a simple command:

`nix run github:karavel-io/cli <regular karavel arguments>`

e.g.

`nix run github:karavel-io/cli render --debug`

Or you can import it in another `flake.nix`, e.g. to add it to a [devShell](https://nixos.wiki/wiki/Development_environment_with_nix-shell):

```nix
{
  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs";
    flake-utils.url = "github:numtide/flake-utils";
    karavel.url = "github:karavel-io/cli";
  };

  outputs = { self, nixpkgs, flake-utils, karavel }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
        karavel-cli = karavel.defaultPackage.${system};
      in {
        devShell = pkgs.mkShell { buildInputs = with pkgs; [
          karavel-cli
          kubectl
          kustomize
        ];
        };
      });
}
```

By appending a refernce at the end of the Flake URL you can select:

- a specific tag/version: `github:karavel-io/cli/v0.4.1`
- a Git branch: `github:karavel-io/cli/main`

More information on references can be found on the [Flakes manual](https://nixos.org/manual/nix/stable/command-ref/new-cli/nix3-flake.html#flake-references).

### Docker

The CLI is packaged in a container image and published on [Quay](https://quay.io/karavel/cli) and [GitHub](https://github.com/karavel-io/cli/pkgs/container/cli).

You can run it like so:

```bash
# Inside a Karavel project directory
$ docker run --rm -v $PWD:/karavel -u (id -u) quay.io/karavel/cli:main render
$ docker run --rm -v $PWD:/karavel -u (id -u) ghcr.io/karavel-io/cli:main render
```

Stable releases are tagged using their semver (`x.y.z`) version, with aliases to the latest patch (`x.y`) and minor (`x`) versions.
This is what you should be using most of the time.  
The `main` tag points to the latest unstable build from the `main` branch. It's useful if you want to try out the latest
features before they are released.

### From sources

Instructions for compiling the tool from source code can be found on [GitHub](https://github.com/karavel-io/cli).

## karavel init

Initialize a new Karavel project

Usage:
  karavel init [WORKDIR] [flags]

Flags:
      --force                Overwrite the config file even if it already exists
      --github-repo string   Override the official GitHub repository containing the tagged Platform releases
  -h, --help                 help for init
  -o, --output-file string   Karavel config file name to create (default "karavel.hcl")
  -v, --version string       Karavel Container Platform version to initialize (default "latest")

Global Flags:
      --colors   Enable colored logs (default true)
  -d, --debug    Output debug logs
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

Fetching latest release version for GitHub repo karavel-io/platform

Writing config file to /tmp/karavel/karavel.hcl
```

To download a specific version of the Karavel Container Platform instead of the latest one, you can pass the `--version`
flag to `karavel init` with the desired version.

```bash
$ cd /tmp/karavel 
$ karavel init --version 2021.1
Initializing new Karavel 2021.1 project at /tmp/karavel


Writing config file to /tmp/karavel/karavel.hcl
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
      --skip-git      Skip the git integration to discover remote repositories for Argo. WARNING: this will render the Argo component inoperable

Global Flags:
      --colors   Enable colored logs (default true)
  -d, --debug    Output debug logs
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
