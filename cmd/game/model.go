package main

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	mapStyle = lipgloss.NewStyle().
		Border(lipgloss.BlockBorder())
)

type model struct {
	h, w   int
	player Entity
}

func (m model) up() (tea.Model, tea.Cmd) {
	if m.player.Pos.Y < m.h {
		m.player.up()
	}

	return m, nil
}

func (m model) down() (tea.Model, tea.Cmd) {
	if m.player.Pos.Y > m.h*-1 {
		m.player.down()
	}

	return m, nil
}

func (m model) left() (tea.Model, tea.Cmd) {
	if m.player.Pos.X > m.w*-1 {
		m.player.left()
	}

	return m, nil
}

func (m model) right() (tea.Model, tea.Cmd) {
	if m.player.Pos.X < m.w {
		m.player.right()
	}

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
	for y := m.h; y > m.h*-1; y-- {
		for x := m.w * -1; x < m.w; x++ {
			if m.player.Pos.at(x, y) {
				fmt.Fprint(b, "X")
			} else {
				fmt.Fprint(b, " ")
			}
		}

		fmt.Fprintln(b)
	}

	world := mapStyle.Render(b.String())

	return strings.Join([]string{
		m.player.Pos.String(),
		world,
	}, "\n")
}
