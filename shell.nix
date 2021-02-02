let
  pkgs = import <nixpkgs> { };
  unstable = import (fetchTarball https://github.com/NixOS/nixpkgs-channels/archive/nixos-unstable.tar.gz) { };
  addlicense = pkgs.callPackage ./.nix/addlicense.nix { };
  jb = pkgs.callPackage ./.nix/jb.nix { inherit pkgs; };
in
pkgs.mkShell {
  buildInputs = with pkgs; [
    nixpkgs-fmt
    go
    kubectl
    kubernetes-helm
    kustomize
    kind
    addlicense
    unstable.velero
    jsonnet
    jb
  ];
}
