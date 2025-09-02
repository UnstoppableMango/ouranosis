package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/unmango/go/cli"
)

func main() {
	p := tea.NewProgram(model{
		h: 10,
		w: 25,
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
