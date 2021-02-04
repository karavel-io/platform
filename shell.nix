let
  pkgs = import <nixpkgs> {};
  openshift = import ./.nix/openshift.nix { pkgs = pkgs; };
  pypkgs = pkgs.python38Packages;
in
pkgs.mkShell {
    buildInputs = with pkgs; [
      mkdocs
      ansible_2_10
      openshift
      pypkgs.kubernetes
      go
    ];
}
