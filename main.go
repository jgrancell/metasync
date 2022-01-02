package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/jgrancell/metasync/app"
	"github.com/jgrancell/metasync/configuration"
)

var (
	configPath string
	debug      bool
	diff       bool
	dryRun     bool
	subcommand string
	verbose    bool

	version string = "0.1.0"
)

func main() {
	os.Exit(Run())
}

func Run() int {

	// We could set the flags normally, but in the event we need different register and deregister flags in the future
	// I've added this as a Flagset for easy duplication between the two commands.
	flagset := flag.NewFlagSet("standard", flag.ExitOnError)
	flagset.StringVar(&configPath, "conf", ".metasync.yml", "The metasync YAML configuration file to read from.")
	flagset.BoolVar(&dryRun, "dryrun", false, "Dry-run outputs the changes that would be made without changing any files.")
	flagset.BoolVar(&verbose, "verbose", false, "Provide verbose output.")
	flagset.BoolVar(&debug, "debug", false, "Provide debugging output.")
	flagset.BoolVar(&diff, "diff", false, "Show diffs for files that require sync.")

	// Detecting our subcommand and parsing CLI flags
	switch os.Args[1] {
	case "sync":
		subcommand = "sync"
		flagset.Parse(os.Args[2:])
	default:
		fmt.Println("The first argument to metasync must be a valid subcommand. These are:")
		fmt.Println("  sync")
		return 1
	}

	// Setting up our configurations and initializing the application
	conf := &configuration.Configuration{
		ConfigurationPath: configPath,
	}
	if err := conf.Load(); err != nil {
		fmt.Println(err.Error())
		return 1
	}

	// Initializing the application
	app := &app.App{
		Configuration: conf,
		Debug:         debug,
		Dryrun:        dryRun,
		ShowDiffs:     diff,
		Subcommand:    subcommand,
		Verbose:       verbose,
		Version:       version,
	}
	return app.Run()
}
