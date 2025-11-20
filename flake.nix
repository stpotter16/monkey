{
    description = "Dev environment for munnin";

    inputs = {
        flake-utils.url = "github:numtide/flake-utils";

        # 1.25.1
        go-nixpkgs.url = "github:NixOS/nixpkgs/bce5fe2bb998488d8e7e7856315f90496723793c";
    };

    outputs = {
        self,
        flake-utils,
        go-nixpkgs,
    } @ inputs:
    flake-utils.lib.eachDefaultSystem (system: let
        gopkg = go-nixpkgs.legacyPackages.${system};
        go = gopkg.go_1_25;
    in {
        devShells.default = gopkg.mkShell {
            packages = [
                gopkg.gotools
                gopkg.gopls
                gopkg.go-outline
                gopkg.gopkgs
                gopkg.gocode-gomod
                gopkg.godef
                gopkg.golint
                go
            ];

            shellHook = ''
                export GOROOT="${go}/share/go"
                go version
            '';
        };
    });
}
