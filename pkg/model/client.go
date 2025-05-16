package model

import tea "github.com/charmbracelet/bubbletea/v2"

type Client struct{}

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
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			cmd = tea.Quit
		}
	}

	return c, cmd
}

func (c Client) View() string {
	return "Test"
}
