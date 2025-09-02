package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/unmango/go/cli"
	"github.com/unstoppablemango/ouranosis/pkg/domain"
	"github.com/unstoppablemango/ouranosis/pkg/tui"
)

type model struct {
	player domain.Player
}

func (m model) up() (tea.Model, tea.Cmd) {
	m.player.Up()
	return m, nil
}

func (m model) down() (tea.Model, tea.Cmd) {
	m.player.Down()
	return m, nil
}

func (m model) left() (tea.Model, tea.Cmd) {
	m.player.Left()
	return m, nil
}

func (m model) right() (tea.Model, tea.Cmd) {
	m.player.Right()
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

	for y := 10; y > -10; y-- {
		for x := -10; x < 10; x++ {
			if m.player.At(x, y) {
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
		player: tui.NewPlayer("Jeff"),
	})
	if _, err := p.Run(); err != nil {
		cli.Fail(err)
	}
}
