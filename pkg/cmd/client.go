package cmd

import (
	tea "github.com/charmbracelet/bubbletea/v2"
	"github.com/spf13/cobra"
	"github.com/unmango/go/cli"
	"github.com/unstoppablemango/ouranosis/pkg/model"
)

func NewClient() *cobra.Command {
	return &cobra.Command{
		Use: "client",
		Run: func(cmd *cobra.Command, args []string) {
			m := model.NewClient()
			prog := tea.NewProgram(m,
				tea.WithContext(cmd.Context()),
				tea.WithAltScreen(),
			)
			if _, err := prog.Run(); err != nil {
				cli.Fail(err)
			}
		},
	}
}
