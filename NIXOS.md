# Dfetch on NixOS

This repository provides a Nix flake for building, running, and installing Dfetch on NixOS and other Linux systems with Nix.

## Use the Flake

Add Dfetch to your flake inputs:

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    dfetch.url = "github:David17c/Dfetch";
  };
}
```

Install the package in your NixOS configuration:

```nix
{ inputs, pkgs, ... }:

{
  environment.systemPackages = [
    inputs.dfetch.packages.${pkgs.stdenv.hostPlatform.system}.default
  ];
}
```

Then rebuild:

```bash
sudo nixos-rebuild switch --flake /path/to/your/flake#your-hostname
```

## Use the NixOS Module

The flake also exposes a small NixOS module that installs Dfetch through `programs.dfetch`.

```nix
{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    dfetch.url = "github:David17c/Dfetch";
  };

  outputs = { self, nixpkgs, dfetch, ... }: {
    nixosConfigurations.your-hostname = nixpkgs.lib.nixosSystem {
      system = "x86_64-linux";
      modules = [
        dfetch.nixosModules.default
        {
          programs.dfetch.enable = true;
        }
      ];
    };
  };
}
```

## Run Without Installing

You can run Dfetch directly from the flake:

```bash
nix run github:David17c/Dfetch
```

## Development Shell

Enter a shell with Go and Git available:

```bash
nix develop github:David17c/Dfetch
```

## Package Counting on Nix

When the `packages` module is enabled in Dfetch's config, Nix package counting checks the standard system and user profile paths and combines their results. Dfetch queries each profile's Nix requisites and filters them with the same package-oriented rules used by Fastfetch, so it is not limited to executable links in `bin`.

- `/run/current-system` for the active NixOS system profile
- `~/.nix-profile` for the user's default Nix profile
- `$XDG_STATE_HOME/nix/profile` or `~/.local/state/nix/profile` for the newer user profile location
- `/etc/profiles/per-user/$USER` for the per-user profile

Missing profile directories are ignored, so Dfetch still works on systems that only have some of these paths.

## Dfetch Configuration

Dfetch reads its runtime configuration from:

```text
~/.config/dfetch/dfetch.conf
```

Include `packages` in the modules block to show package information:

```text
modules {
    userinfo
    os
    kernel
    packages
    memory
    disk
}
```

For general configuration and usage, see the main [README](README.md).
