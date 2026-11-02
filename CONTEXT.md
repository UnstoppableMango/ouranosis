# Ouranosis

A world simulation game that starts with one character and grows outward, built on the `unmango/game` framework.

## Language

**Slice**:
One increment of the build: a playable loop with a "done when" that is shipped as a unit.
_Avoid_: milestone, phase, release

**Stage**:
A layer of the world in the vision, from the character outward: character, party, settlement, army, city, world.
One slice can touch several stages, and a stage is usually built across several slices.
_Avoid_: tier, level, era

**Sub-game**:
An activity that reads and writes the character, addressed by a path on the framework.
_Avoid_: minigame, mode

## Relationships

- A **Slice** delivers playable work in one or more **Stages**
- A **Stage** contains many **Sub-games**

## Example dialogue

> **Dev:** "Is recruitment the next **slice**?"
> **Designer:** "It's a **slice** in the party **stage**, but the slice also needs a town to recruit in, so it touches the settlement **stage** too."
