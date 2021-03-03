# Karavel CLI

TODO

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
