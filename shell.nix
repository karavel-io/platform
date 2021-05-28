let
  pkgs = import <nixpkgs> { };
  addlicense = pkgs.callPackage ./.nix/addlicense.nix { };
in
pkgs.mkShell {
  buildInputs = with pkgs; [
    nixpkgs-fmt
    go_1_16
    kubectl
    kubernetes-helm
    kustomize
    pkgs.unstable.kind
    addlicense
    pkgs.unstable.velero
    bats
  ];
}
