# AGENTS.md

Guidance for AI agents working in this repository.

## Overview

Ouranosis is a world simulation game built on the `github.com/unmango/game` framework.
It begins as a role playing game around one character and grows outward: party, army, city, world.
Every activity is a sub-game that feeds the same character, and the framework keeps the numbers consistent across them.

The design lives in `docs/design/` and the decisions in `docs/decisions/`.
Read `docs/design/vision.md` for what the game is and `docs/design/first-slice.md` for what is being built.

## Commands

All development tasks go through `make`.

```sh
make            # go build into bin/ouranosis
make run        # build and start the TUI
make test       # go tool ginkgo run -r
make check      # go vet
make fmt        # gofmt
make tidy       # go mod tidy
```

The TUI expects a framework server at `localhost:8080`.
Start one from a checkout of `unmango/game` with `go run ./cmd/game`.

## Architecture

```
cmd/ouranosis/      the TUI binary
pkg/game/           client for the framework's ConnectRPC services
pkg/character/      character model and the paths it uses on the framework
pkg/training/       the stat training sub-game
docs/design/        design documents
docs/decisions/     architecture decision records
```

Abstractions belong in the framework; content and rules belong here.
Before adding a numeric helper, check whether `unmango/game` already answers the question, and if it does not, decide whether the addition is math (framework) or meaning (here).
See ADR 0003.

Every path this game uses on the framework is a public contract with its players, since renaming a path reseeds it.
Paths are declared as constants in `pkg/character` and nowhere else.

Tests use Ginkgo and Gomega.
