# 0001. Ouranosis is built on unmango/game

## Status

Accepted

## Context

The game needs prices, growth rates, rarity, and offline progress that stay consistent across many sub-games and grow without bound.
Writing that math inside the game would couple every sub-game to every other.

## Decision

All numeric progression comes from the `github.com/unmango/game` framework over ConnectRPC.
The game imports the generated Go client from the framework module and holds no curve math of its own.
Needs that arise here drive the framework's priorities.

## Consequences

- The game runs against a local framework server during development.
- Every sub-game is a set of paths and curve parameters, which keeps them uniform.
- A framework change can be breaking for the game; the two are developed together.
