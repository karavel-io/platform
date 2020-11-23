{ pkgs, ...}:
let
pypkgs = pkgs.python38Packages;
string_utils = pypkgs.buildPythonPackage rec {
    name = "python-string-utils-1.0.0";

    src = pkgs.fetchurl {
      url = "https://files.pythonhosted.org/packages/10/91/8c883b83c7d039ca7e6c8f8a7e154a27fdeddd98d14c10c5ee8fe425b6c0/${name}.tar.gz";
      sha256 = "dcf9060b03f07647c0a603408dc8b03f807f3b54a05c6e19eb14460256fac0cb";
    };

    meta = {
      description = "Utility functions for strings validation and manipulation.";
    };
  };
in
pypkgs.buildPythonPackage rec {
    name = "openshift-0.11.2";

    buildInputs = [
        string_utils
        pypkgs.ruamel_yaml
        pypkgs.six
        pypkgs.jinja2
        pypkgs.kubernetes
    ];
    src = pkgs.fetchurl {
      url = "https://files.pythonhosted.org/packages/2a/f2/978b34965425fa737464082ad96d46646ada88fb94f6f84ee2f8581df305/${name}.tar.gz";
      sha256 = "110b0d3c84a83500f0fd150ab26dee29615157e6659bf72808788aa79fc17afc";
    };

    meta = {
      description = "OpenShift python client";
    };
}
