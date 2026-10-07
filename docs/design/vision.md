# Vision

Ouranosis is a world simulation that starts with one character and never stops growing.

## Premise

The player wakes into a world with no instructions and a few things to do.
What they do first shapes who their character becomes, and the story is set up by the player's choices rather than told to them.
From that start the game widens: companions, a party, a settlement, an army, a city, a world.
Every stage is a sub-game that reads and writes the same character.

## Pillars

1. **Agency from the first minute.** The opening is open, in the manner of Skyrim's first hour, and the story assembles itself from what the player chooses.
2. **Every action counts.** Time spent on any activity leaves a lasting mark on the character.
   Nothing is wasted and nothing is lost.
   This is the framework's one hard rule, and the game is built to honor it.
3. **One character, many genres.** The same character, party, and equipment carry into every sub-game, scaled by the framework so that each activity is challenging at the player's level.
4. **Recursion.** A stat, an item, a building, or a companion can be opened as its own sub-game.
   Playing that sub-game levels the thing it belongs to.

## Inspirations, by mechanic

| Mechanic | Source |
| --- | --- |
| Open start, story from choice | Skyrim |
| Stats that grow by use | Bethesda role playing games |
| Character creation with classes and reincarnation | Disgaea |
| Grid tactics encounters | Disgaea |
| Party building, recruitment, town to overworld to dungeon flow | Dragon Quest V |
| Love interest and marriage | Dragon Quest V |
| First person combat with the same party and equipment | Shooters generally |
| City and army building as sub-games | Strategy games generally |

## Structure

The root game is the character's life.
Sub-games hang off it as paths on the framework:

```
character/                the root
character/<stat>/         training sub-games, one per stat
character/equipment/      items, each openable as its own sub-game
party/                    companions and recruitment
encounter/tactics/        grid battles
encounter/first-person/   real time battles
settlement/               city building
settlement/<business>/    businesses inside the city, each a sub-game
```

Each sub-game declares which root paths it advances and by how much.
The framework's curves scale each one to the character's level.

## Multiplayer

Sub-games with a competitive form (tactics, first person, city trade) can exchange values with other players through the framework's federation.
Battles can resolve offline: both sides compute the same outcome from the same inputs and reconcile on the next connection.
This is a later stage and does not shape the first slice.

## Non-goals for now

- Graphics beyond a terminal.
- Real time combat.
- Multiplayer.
- Story content.
