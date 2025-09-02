package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/unmango/go/cli"
)

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
	X, Y int64
}

func (p Point) at(x, y int64) bool {
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

type model struct {
	player Entity
}

func (m model) up() (tea.Model, tea.Cmd) {
	m.player.up()
	return m, nil
}

func (m model) down() (tea.Model, tea.Cmd) {
	m.player.down()
	return m, nil
}

func (m model) left() (tea.Model, tea.Cmd) {
	m.player.left()
	return m, nil
}

func (m model) right() (tea.Model, tea.Cmd) {
	m.player.right()
	return m, nil
}

// Init implements tea.Model.
func (m model) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up":
			return m.up()
		case "down":
			return m.down()
		case "left":
			return m.left()
		case "right":
			return m.right()
		}
	}

	return m, nil
}

// View implements tea.Model.
func (m model) View() string {
	b := &strings.Builder{}
	// fmt.Fprintf(b, "%s\n", m.player)

	for y := 10; y > -10; y-- {
		for x := -10; x < 10; x++ {
			if m.player.Pos.at(int64(x), int64(y)) {
				fmt.Fprint(b, "X")
			} else {
				fmt.Fprint(b, " ")
			}
		}

		fmt.Fprintln(b)
	}

	return b.String()
}

func main() {
	p := tea.NewProgram(model{
		player: Entity{
			Character: Character{
				Name: "Jeff",
			},
			Pos: Point{0, 0},
		},
	})
	if _, err := p.Run(); err != nil {
		cli.Fail(err)
	}
}
