let
  pkgs = import <nixpkgs> {};
  linkerd = pkgs.callPackage .nix/linkerd.nix { inherit pkgs; };
  unstable = import (fetchTarball https://github.com/NixOS/nixpkgs-channels/archive/nixos-unstable.tar.gz) { };
  addlicense = pkgs.callPackage ./.nix/addlicense.nix { };
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
    linkerd
  ];
}
