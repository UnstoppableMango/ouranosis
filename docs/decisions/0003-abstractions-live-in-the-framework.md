# 0003. Abstractions live in the framework, content lives here

## Status

Accepted

## Context

As sub-games multiply, helpers that several of them share will appear.
Some belong to the framework and some are rules of this game.
Without a test, the framework either bloats or the game grows a second framework inside it.

## Decision

A helper moves to `unmango/game` when it is math or identity: a curve, a derivation, a rate, an exchange.
It stays here when it gives meaning to a number: what a stat is, what an upgrade does, what a class modifies.
When a move would make the framework know about game content, it stays here even if several sub-games share it.

## Consequences

- The framework stays a thin spine.
- Shared game content lives in `pkg/` packages with no terminal dependency.
- Each move to the framework is a framework PR first, then a game PR that adopts it.
