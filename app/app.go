package app

import (
	"fmt"

	"github.com/jgrancell/metasync/configuration"
)

type App struct {
	Configuration *configuration.Configuration
	Dryrun        bool
	Subcommand    string
	Verbose       bool
	Version       string
}

func (a *App) Run() int {
	// Running our subcommand
	if a.Subcommand == "sync" {
		// TODO: add app.Sync() or similar here
		return 0
	} else {
		fmt.Println("Specified subcommand", a.Subcommand, "is not valid.")
		return 1
	}
}
