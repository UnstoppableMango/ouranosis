# First slice

The first playable loop: create a character, open one stat, train it in a sub-game, return, and see the stat raised.
It exercises seed derivation, curves, sub-game nesting, and offline progress in one loop.

## Screens

1. **Character creation.**
   Name and class.
   Classes are a fixed list (warrior, mage, rogue).
   Base stats are rolled from `Derive("character/<stat>")` with class modifiers, so the same server always rolls the same character for the same class.
2. **Character sheet.**
   Name, class, and each stat with its level.
   Selecting a stat opens its training sub-game.
3. **Training.**
   An incremental for one stat.
   Experience accrues at a rate per second.
   Upgrades multiply the rate and are priced by an exponential curve from the framework.
   The player leaves at any time.

## Numbers

All numbers come from the framework.

| Question | Framework call |
| --- | --- |
| Base value of a stat | `Derive("character/<stat>")` reduced to a range by the game |
| Cost of the nth upgrade | `Evaluate(exponential{base, growth}, n)` |
| How many upgrades a balance buys | `Invert(curve, owned, balance)` |
| Experience gained while away | `Advance(rate_curve, upgrades, elapsed)` |
| Experience to reach level L | `Evaluate(polynomial{scale, degree}, L)` |

Curve parameters are constants in `pkg/training` for the first slice.
Moving them to a tuning file is a later step.

## Paths

```
character/<stat>              base roll
character/<stat>/training     the sub-game root
```

Stat names for the first slice: `strength`, `agility`, `mind`.

## Persistence

Local state (character, per stat experience, upgrades owned, last seen time) is a JSON file next to the framework's data.
On open, the game asks the framework how much experience elapsed since last seen and applies it.
The framework ledger replaces this file once it exists; the JSON file is a stand-in with the same shape.

## Done when

The player can create a character, enter a stat, buy an upgrade, quit, reopen, and see both the offline experience and the higher stat level.
