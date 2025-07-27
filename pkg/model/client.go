package model

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/charmbracelet/lipgloss"
)

var screen = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("201"))

var title = lipgloss.NewStyle().
	Align(lipgloss.Center).
	Bold(true)

type Client struct {
	Height, Width int
	Game          *Game
}

func NewClient() Client {
	return Client{}
}

// Init implements tea.Model.
func (c Client) Init() tea.Cmd {
	return nil
}

// Update implements tea.Model.
func (c Client) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.Height = msg.Height
		c.Width = msg.Width
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			cmd = tea.Quit
		}
	}

	return c, cmd
}

func (c Client) View() string {
	height, width := c.screenSize()
	s := screen.Width(width).Height(height)

	t := title.Width(width).Height(height)
	title := t.Render("Ouranosis")

	return s.Render(title)
}

func (c Client) screenSize() (height, width int) {
	hz := screen.GetHorizontalBorderSize()
	vz := screen.GetVerticalBorderSize()

	return c.Height - vz, c.Width - hz
}
