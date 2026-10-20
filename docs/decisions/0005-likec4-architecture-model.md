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

This repository has no Nix, so `likec4` comes from npm.
`docs/architecture/package.json` pins an exact version with a lockfile, which keeps it reproducible and lets Renovate bump it with the other dependencies.
The version matches the one `unmango/pkgs` packages for the framework.
The `likec4/actions` action is not used, so the version lives in one place for both local runs and CI.

The site is built with a relative base and hash history, so one build works at the Pages URL, a custom domain, or opened from disk.
The `pages` workflow builds it and deploys it to GitHub Pages on every push to `main`, and CI validates and builds it on every pull request.
Elements link to their source directories and the ADRs that shaped them, so the site is a map into the repository.

## Consequences

- The model changes alongside the code it describes; adding a package, client, or store means adding it to the model.
- Working on the model needs Node.js; `make arch`, `make arch-check`, and `make arch-build` install the pinned CLI on first use. Go work does not.
- Rendered output is not committed.
- Publishing needs Pages enabled with "GitHub Actions" as the source in the repository settings.
- The two models are separate projects. A combined site that imports the framework's model is possible with a LikeC4 workspace but is deferred until the slice drawn here outgrows it.
