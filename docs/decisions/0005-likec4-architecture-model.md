# 0005. Architecture is modeled in LikeC4

## Status

Accepted

## Context

`unmango/game` models its architecture in LikeC4 (its ADR 0007) and publishes it to GitHub Pages.
That model draws ouranosis only where it touches the framework, and leaves the game's own detail to the game.
As sub-games, clients, and atproto records arrive, the shape of the game should be reviewable as text next to the code, the same way the framework's is.

## Decision

The architecture is a LikeC4 model in `docs/architecture/`, laid out like the framework's: `spec.c4` for kinds and tags, `model.c4` for elements and relationships, `views.c4` for views.
Designed but unimplemented parts carry `#planned` and render dashed.
Framework parts appear only where the game touches them and carry `#framework`; their detail stays in the framework's model, which they link to.

The site is built by Nix with `lib.likec4` from `github:UnstoppableMango/a2b`, the same builder the framework uses, so this repository maintains no Node.js tooling.
A flake exists only for this; the Go build stays plain Go and needs no Nix.
Its `a2b` input follows `unmango-pkgs`, so the site, the check, and the dev shell share one pinned `likec4`, and the version moves with `nix flake update`.
The `likec4/actions` action is not used because it installs `likec4` from npm outside that pin.

It uses a relative base and hash history, so one build works at the Pages URL, a custom domain, or opened from disk.
The `pages` workflow builds it with `nix build .#architecture` and deploys it to GitHub Pages on every push to `main`.
`nix flake check` validates the model and builds the site, and CI runs it on every pull request, so a model that validates but fails to render still fails CI.
Elements link to their source directories and the ADRs that shaped them, so the site is a map into the repository.

## Consequences

- The model changes alongside the code it describes; adding a package, client, or store means adding it to the model.
- Working on the model needs Nix: `make arch` previews it from the dev shell (`nix develop`), `make arch-check` runs `nix flake check`, and `make arch-build` builds the site into `result/`. Go work does not.
- Rendered output is not committed.
- Publishing needs Pages enabled with "GitHub Actions" as the source in the repository settings.
- The two models are separate projects. A combined site that imports the framework's model is possible with a LikeC4 workspace but is deferred until the slice drawn here outgrows it.
