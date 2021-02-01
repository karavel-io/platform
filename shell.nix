let
  pkgs = import <nixpkgs> {};
  openshift = pkgs.callPackage .nix/openshift.nix { inherit pkgs; };
  linkerd = pkgs.callPackage .nix/linkerd.nix { inherit pkgs; };
  pypkgs = pkgs.python38Packages;
in
pkgs.mkShell {
    buildInputs = with pkgs; [
      mkdocs
      ansible_2_10
      openshift
      pypkgs.kubernetes
      linkerd
    ];
}
