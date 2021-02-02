{ pkgs, stdenv, buildGoModule, fetchFromGitHub }:

buildGoModule rec {
    pname = "jsonnet-bundler";
    version = "0.4.0";

    src = fetchFromGitHub {
        owner = "jsonnet-bundler";
        repo = "jsonnet-bundler";
        rev = "v${version}";
        sha256 = "0pk6nf8r0wy7lnsnzyjd3vgq4b2kb3zl0xxn01ahpaqgmwpzajlk";
    };

    vendorSha256 = null;

    subPackages = [ "cmd/jb" ];
    meta = with stdenv.lib; {
      description = "A jsonnet package manager";
      homepage = "https://github.com/jsonnet-bundler/jsonnet-bundler";
      license = licenses.asl20;
      maintainers = with maintainers; [ matteojoliveau ];
      platforms = platforms.unix;
    };
}
