{
  description = "A world simulation game built on unmango/game";

  inputs = {
    nixpkgs.url = "github:nixos/nixpkgs?ref=nixos-unstable";
    systems.url = "github:UnstoppableMango/nix-systems";

    flake-parts = {
      url = "github:hercules-ci/flake-parts";
      inputs.nixpkgs-lib.follows = "nixpkgs";
    };

    # Not used here; declared so unmango-pkgs and a2b share one copy.
    treefmt-nix = {
      url = "github:numtide/treefmt-nix";
      inputs.nixpkgs.follows = "nixpkgs";
    };

    # Not used here; declared so unmango-pkgs shares one copy.
    gomod2nix = {
      url = "github:nix-community/gomod2nix";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.flake-utils.inputs.systems.follows = "systems";
    };

    unmango-pkgs = {
      url = "github:unmango/pkgs";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.systems.follows = "systems";
      inputs.flake-parts.follows = "flake-parts";
      inputs.gomod2nix.follows = "gomod2nix";
      inputs.treefmt-nix.follows = "treefmt-nix";
    };

    a2b = {
      url = "github:UnstoppableMango/a2b";
      inputs.nixpkgs.follows = "nixpkgs";
      inputs.systems.follows = "systems";
      inputs.flake-parts.follows = "flake-parts";
      inputs.treefmt-nix.follows = "treefmt-nix";
      inputs.mangopkgs.follows = "unmango-pkgs";
    };
  };

  outputs =
    inputs@{ flake-parts, ... }:
    flake-parts.lib.mkFlake { inherit inputs; } {
      systems = import inputs.systems;

      imports = with inputs; [
        systems.flakeModule or { }
      ];

      # The Go build stays plain Go; the flake only builds the architecture site.
      perSystem =
        {
          pkgs,
          self',
          inputs',
          ...
        }:
        let
          inherit (inputs'.unmango-pkgs.packages) likec4;
          arch = inputs'.a2b.legacyPackages.lib.likec4;
          archSrc = ./docs/architecture;
        in
        {
          # Relative base and hash history let one build serve from any URL prefix.
          packages.architecture = arch.build {
            name = "architecture";
            src = archSrc;
            base = "./";
            useHashHistory = true;
            title = "ouranosis architecture";
          };

          devShells.default = pkgs.mkShellNoCC {
            packages = [
              likec4
              pkgs.gnumake
            ];
          };

          checks.architecture = arch.validate {
            name = "architecture";
            src = archSrc;
          };
          checks.architecture-site = self'.packages.architecture;
        };
    };
}
