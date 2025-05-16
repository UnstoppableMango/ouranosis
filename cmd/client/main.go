package main

import (
	"github.com/unmango/go/cli"
	"github.com/unstoppablemango/ouranosis/pkg/cmd"
)

func main() {
	cmd := cmd.NewClient()
	if err := cmd.Execute(); err != nil {
		cli.Fail(err)
	}
}
