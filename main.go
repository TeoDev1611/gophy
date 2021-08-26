package main

import (
	"errors"
	"fmt"

	"github.com/TeoDev1611/gophy/cmds"
	"github.com/tbuchaillot/icli"
)

func main() {
	cli := icli.NewCLI()
	cli.SetErrorColor(icli.PURPLE)

	cli.SetWelcomeMessage(fmt.Sprintf("%v Welcome to Gophy a attempt to make a simple shell!%v", icli.GREEN, icli.RESET))

	versionCmd := &icli.BasicCommand{
		Name:        "version",
		Description: "Show the version of Gophy",
		Usage:       "version",
		Fn:          Version,
	}

	lsCmd := &icli.BasicCommand{
		Name:        "ls",
		Description: "List all directories in the current directory",
		Usage:       "ls",
		Fn:          cmds.ListFiles,
	}

	cli.AddCmds(versionCmd, lsCmd)

	cli.Run()
}

func Version(args ...string) error {
	if len(args) <= 0 {
		fmt.Println("Version 0.1.0")
		return nil
	} else {
		return errors.New("Not arguments needed")
	}
}
