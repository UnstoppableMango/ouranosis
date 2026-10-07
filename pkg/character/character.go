// Package character holds the player character and the framework paths it
// owns. Paths are a contract with players: renaming one reseeds it.
package character

import (
	"context"
	"fmt"
	"time"

	"github.com/unmango/game/num"

	"github.com/unstoppablemango/ouranosis/pkg/game"
)

// Class is a starting archetype.
type Class string

// The classes available at creation.
const (
	Warrior Class = "warrior"
	Mage    Class = "mage"
	Rogue   Class = "rogue"
)

// Classes lists every class in display order.
var Classes = []Class{Warrior, Mage, Rogue}

// Stat is a trainable attribute.
type Stat string

// The stats in the first slice.
const (
	Strength Stat = "strength"
	Agility  Stat = "agility"
	Mind     Stat = "mind"
)

// Stats lists every stat in display order.
var Stats = []Stat{Strength, Agility, Mind}

// Path is the framework path for the stat's base roll.
func (s Stat) Path() string {
	return "character/" + string(s)
}

// TrainingPath is the framework path for the stat's training sub-game.
func (s Stat) TrainingPath() string {
	return s.Path() + "/training"
}

var modifiers = map[Class]map[Stat]int64{
	Warrior: {Strength: 3, Agility: 1, Mind: -1},
	Mage:    {Strength: -1, Agility: 1, Mind: 3},
	Rogue:   {Strength: 1, Agility: 3, Mind: -1},
}

// Roll turns a sub-seed into a base stat value between 4 and 13 depending
// on class.
func Roll(sub uint64, class Class, stat Stat) int64 {
	return 5 + int64(sub%6) + modifiers[class][stat]
}

// Progress is what training a stat has produced.
// Earned never decreases; Spent only grows as upgrades are bought.
type Progress struct {
	Earned   num.Number `json:"earned"`
	Spent    num.Number `json:"spent"`
	Upgrades int64      `json:"upgrades"`
	LastSeen time.Time  `json:"last_seen"`
}

// Balance is experience available to spend.
func (p Progress) Balance() num.Number {
	return p.Earned.Sub(p.Spent)
}

// Character is the player.
type Character struct {
	Name     string             `json:"name"`
	Class    Class              `json:"class"`
	Base     map[Stat]int64     `json:"base"`
	Progress map[Stat]*Progress `json:"progress"`
}

// Create rolls a character's base stats from the framework's seed.
func Create(ctx context.Context, calc game.Calculator, name string, class Class, now time.Time) (*Character, error) {
	if _, ok := modifiers[class]; !ok {
		return nil, fmt.Errorf("character: unknown class %q", class)
	}
	c := &Character{
		Name:     name,
		Class:    class,
		Base:     make(map[Stat]int64, len(Stats)),
		Progress: make(map[Stat]*Progress, len(Stats)),
	}
	for _, s := range Stats {
		sub, err := calc.Derive(ctx, s.Path())
		if err != nil {
			return nil, err
		}
		c.Base[s] = Roll(sub, class, s)
		c.Progress[s] = &Progress{LastSeen: now}
	}
	return c, nil
}
