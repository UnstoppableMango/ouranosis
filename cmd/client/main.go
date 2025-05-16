package main

import (
	"github.com/unmango/go/cli"
	"github.com/unstoppablemango/ouranosis/pkg/cmd"
	"github.com/unstoppablemango/ouranosis/pkg/cmd/client"
)

func main() {
	cmd := cmd.NewClient()
	cmd.AddCommand(client.NewTextCommand())

	if err := cmd.Execute(); err != nil {
		cli.Fail(err)
	}
}
