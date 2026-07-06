{
  description = "A lightweight system information tool focused on clean output";

  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = { self, nixpkgs, flake-utils }:
    flake-utils.lib.eachDefaultSystem (system:
      let
        pkgs = nixpkgs.legacyPackages.${system};
      in
      {
        packages.default = pkgs.buildGoModule {
          pname = "dfetch";
          version = "1.0.0";
          src = self;

          vendorHash = null;

          installPhase = ''
            mkdir -p $out/bin
            cp dfetch $out/bin/
          '';

          meta = with pkgs.lib; {
            description = "A lightweight system information tool focused on clean output";
            homepage = "https://github.com/crispdark/Dfetch";
            license = licenses.mit;
            maintainers = [];
            platforms = platforms.linux;
          };
        };

        devShells.default = pkgs.mkShell {
          buildInputs = with pkgs; [
            go
            git
          ];
        };
      }
    );
}
