// +build go1.8

package main

import (
	"os"

	"github.com/shawncatz/bub/commands"
	"github.com/shawncatz/bub/github"
	"github.com/shawncatz/bub/ui"
)

func main() {
	defer github.CaptureCrash()

	err := commands.CmdRunner.Execute()
	if !err.Ran {
		ui.Errorln(err.Error())
	}
	os.Exit(err.ExitCode)
}
