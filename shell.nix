let
  pkgs = import <nixpkgs> { };
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
  ];
}
