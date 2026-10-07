# 0002. The first client is a terminal UI

## Status

Accepted

## Context

The first slice needs a client that can be built in the same repo, in the same language, with no build pipeline, so that the loop is playable early.
A web client and a graphical client are both wanted later.

## Decision

The first client is a bubbletea terminal application in `cmd/ouranosis`.
Game logic lives in packages under `pkg/` with no dependency on the terminal, so a later client reuses it.

## Consequences

- No graphics until a second client exists.
- Screens are bubbletea models; state and rules are plain Go.
- The web client is a stack layer after the first slice works.
