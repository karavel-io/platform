{ lib, buildGoModule, fetchFromGitHub }:

buildGoModule rec {
    pname = "fetch";
    version = "v0.4.2";

    src = fetchFromGitHub {
        owner = "gruntwork-io";
        repo = pname;
        rev = version;
        sha256 = "1i9iplmnccxcxd63412f7mh1yp9zgi00d5svyi8hmvjblp84sm61";
    };

    vendorSha256 = "1n88r4isrwgji9dh3s9fx903dw9cy9sv2azyrkrp0fs7an38b5sp";
    doCheck = false;

    meta = with lib; {
        description = "Download files, folders, and release assets from a specific git commit, branch, or tag of public and private GitHub repos.";
        homepage = "https://github.com/gruntwork-io/fetch";
        license = licenses.mit;
        maintainers = with maintainers; [ matteojoliveau ];
        platforms = platforms.unix;
    };
}
