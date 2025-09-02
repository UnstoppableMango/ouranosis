package tui

import "github.com/unstoppablemango/ouranosis/pkg/domain"

type Position struct {
	X, Y int
}

func (p Position) At(x, y int) bool {
	return p.X == x && p.Y == y
}

func (p *Position) Up() {
	p.Y++
}

func (p *Position) Down() {
	p.Y--
}

func (p *Position) Left() {
	p.X--
}

func (p *Position) Right() {
	p.X++
}

type Player struct {
	Position
	Name string
}

func NewPlayer(name string) domain.Player {
	return &Player{
		Name:     name,
		Position: Position{0, 0},
	}
}
