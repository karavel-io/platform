{ lib, buildGoModule, fetchFromGitHub }:

buildGoModule rec {
    pname = "yq";
    version = "v4.9.8";

    src = fetchFromGitHub {
        owner = "mikefarah";
        repo = pname;
        rev = version;
        sha256 = "179fx3yd0lj8cfx6lz43hmn7l3a5cg0q8x58zvqnszr0hmmdpsyw";
    };

    vendorSha256 = "11rh6dc4vgiqrxvi8blhy7rvajybxvcwafj0q0dyb04d1wsvcpjz";
    doCheck = false;

    meta = with lib; {
        description = "yq is a portable command-line YAML processor.";
        homepage = "https://github.com/mikefarah/yq";
        license = licenses.mit;
        maintainers = with maintainers; [ matteojoliveau ];
        platforms = platforms.unix;
    };
}
