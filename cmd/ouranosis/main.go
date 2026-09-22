// Command ouranosis is the terminal client for the first slice.
package main

import (
	"errors"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/unstoppablemango/ouranosis/pkg/character"
	"github.com/unstoppablemango/ouranosis/pkg/game"
	"github.com/unstoppablemango/ouranosis/pkg/save"
)

func main() {
	server := flag.String("server", game.DefaultURL, "framework server URL")
	savePath := flag.String("save", save.DefaultPath(), "save file")
	flag.Parse()

	var ch *character.Character
	loaded, err := save.Load(*savePath)
	switch {
	case err == nil:
		ch = loaded
	case errors.Is(err, fs.ErrNotExist):
	default:
		log.Fatal(err)
	}

	m := newModel(game.New(*server), *savePath, ch)
	if _, err := tea.NewProgram(m, tea.WithAltScreen()).Run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
