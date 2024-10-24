{
  description = "CLI tool to manage Ablaze-MIRAI/Wiki ";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixpkgs-unstable";
    treefmt-nix.url = "github:numtide/treefmt-nix";
    flake-parts.url = "github:hercules-ci/flake-parts";
    systems.url = "github:nix-systems/default";
  };

  outputs =
    inputs@{
      self,
      systems,
      nixpkgs,
      treefmt-nix,
      flake-parts,
    }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      imports = [ treefmt-nix.flakeModule ];
      systems = import inputs.systems;

      perSystem =
        {
          config,
          pkgs,
          system,
          ...
        }:
        let
          stdenv = pkgs.stdenv;

          executable = pkgs.buildGo123Module rec {
            name = "wikitool";
            src = pkgs.fetchFromGitHub {
              owner = "Ablaze-MIRAI";
              repo = "wikitool";
              rev = "main";
              sha256 = "sha256-N1dt/bvKUzpFSMI0e/H3l9zkG01cxyXBoAfxd0950BA=";
            };
            buildInputs = [ pkgs.go ];
            vendorHash = "sha256-OosSVmEzi3cfv9puM9d7+haXmNltjaC/LYKuR6O3KgQ=";
            doCheck = false;
          };
        in
        rec {
          treefmt = {
            projectRootFile = "flake.nix";
            programs = {
              nixfmt.enable = true;
              gofmt.enable = true;
            };

            settings.formatter = { };
          };

          devShells.default = pkgs.mkShell {
            packages = with pkgs; [
              nil
            ];
          };

          packages.default = executable;
        };
    };
}
