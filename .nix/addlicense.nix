{ lib, buildGoModule, fetchFromGitHub }:

buildGoModule rec {
    pname = "addlicense";
    version = "a0294312aa76d31c0bd7e49083d88a2a04d9b3d1";

    src = fetchFromGitHub {
        owner = "google";
        repo = pname;
        rev = version;
        sha256 = "0j2fh9z0yljycxi2x3b9ydcr4a7qg49nw1i9gfvgwk9bz8p7dpw3";
    };

    vendorSha256 = "0cq017iv51vi6m9b1ifwkl522h0yc40cva0jxpx5xywpx1ymxihc";

    meta = with lib; {
        description = "A program which ensures source code files have copyright license headers by scanning directory patterns recursively.";
        homepage = "https://github.com/google/addlicense";
        license = licenses.asl20;
        maintainers = with maintainers; [ matteojoliveau ];
        platforms = platforms.unix;
    };
}