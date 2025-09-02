package main

import "fmt"

type Stat struct {
	Value int64
}

type Character struct {
	Name    string
	Attack  Stat
	Defense Stat
	Speed   Stat
	Luck    Stat
}

func (c Character) String() string {
	return c.Name
}

type Point struct {
	X, Y int
}

func (p Point) at(x, y int) bool {
	return p.X == x && p.Y == y
}

func (p Point) String() string {
	return fmt.Sprintf("X: %d, Y: %d", p.X, p.Y)
}

type Entity struct {
	Character
	Pos Point
}

func (e Entity) String() string {
	return fmt.Sprintf("%s - %s", e.Pos, e.Character)
}

func (e *Entity) up() {
	e.Pos.Y++
}

func (e *Entity) down() {
	e.Pos.Y--
}

func (e *Entity) left() {
	e.Pos.X--
}

func (e *Entity) right() {
	e.Pos.X++
}
