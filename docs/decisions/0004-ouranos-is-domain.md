# 0004. ouranos.is names the game on the AT Protocol

## Status

Accepted

## Context

`unmango/game` ADR 0007 lets a player bind an atproto account to their framework server.
Once players have accounts, the game will publish its own records (characters, and later parties and settlements) and may want to give players readable handles.
Both need a domain: a lexicon's NSID is the reverse of a domain the publisher controls, and a handle is a domain name that resolves to an account.
The framework's lexicons use `dev.unmango.game`, which describes the framework, not this game.

## Decision

The game owns `ouranos.is` and uses it for everything it publishes on atproto.

- Game lexicons use the `is.ouranos` authority, for example `is.ouranos.character`.
  The `_lexicon` DNS records live on `ouranos.is`.
- The game may offer players handles under the domain, such as `name.ouranos.is`.
- Framework records stay under `dev.unmango.game`, per ADR 0003: an atproto record that is math or identity belongs to the framework, and one that gives a number meaning belongs here.

## Consequences

- An NSID is a public contract with players in the same way a path is, since renaming one orphans every record written under it. NSIDs are declared as constants beside the path constants, and nowhere else.
- Losing the domain would orphan both the lexicons and any handles issued under it, so it stays on auto-renew.
- Handles are optional and come later; nothing in the first slice depends on them.
