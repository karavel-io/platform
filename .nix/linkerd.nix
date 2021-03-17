{ pkgs, stdenv, fetchurl }:

let
  name = "linkerd";
  version = "2.10.0";
in
stdenv.mkDerivation {
  name = "${name}";

  src = fetchurl {
    url = "https://github.com/linkerd/linkerd2/releases/download/stable-${version}/linkerd2-cli-stable-${version}-linux-amd64";
    sha256 = "0qvrj373501wp81mzy23hi9y15ccmfybi1z52gddspzpp9ga5k9v";
    executable = true;
  };

  phases = [ "installPhase" "patchPhase" ];

  installPhase = ''
    mkdir -p $out/bin
    cp $src $out/bin/linkerd
  '';

  meta = {
    description = "Ultralight, security-first service mesh for Kubernetes";
    longDescription = ''
      Ultralight, security-first service mesh for Kubernetes.
    '';
    homepage = "https://linkerd.io";
    license = "Apache-2.0";
    platforms = with stdenv.lib.platforms; linux;
    maintainers = [
      stdenv.lib.maintainers.matteojoliveau
    ];
  };
}

