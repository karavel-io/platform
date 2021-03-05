# Karavel CLI

TODO

## Install

TODO: binary install

### Docker

The CLI is packaged in a container image and published on [Quay](https://quay.io/repository/karavel/cli).

You can run it like so:

```bash
# Inside a Karavel project directory
$ docker run --rm -v $PWD:/karavel -u (id -u) quay.io/karavel/cli:edge render
```

The `stable` tag points the latest stable build from a tagged release. This is what you should be using most of the time.  
The `edge` tag points to the latest unstable build from the `master` branch. It's useful if you want to try out the latest
features before they are released.

## Requirements

- Go 1.15+
- make

For [Nix] or [NixOS] users, the provided [shell.nix](../shell.nix) already configures the required tooling.

## Build

`make` outputs the `karavel` executable in the `bin` folder
`make install` installs the executable in the PATH. Install location can be changed by passing the `INSTALL_PATH` variable:
`make INSTALL_PATH=/path/to/karavel install`

[Nix]: https://nixos.org/explore.html
[NixOS]: https://nixos.org
