package app

import (
	"os"
	"testing"

	"github.com/jgrancell/metasync/configuration"
)

func TestSyncSubcommand(t *testing.T) {
	pwd, _ := os.Getwd()

	os.Chdir("../testdata")
	c := &configuration.Configuration{
		ConfigurationPath: ".metasync.yml",
	}
	c.Load()
	os.Chdir(pwd)

	a := &App{
		Configuration: c,
		Dryrun:        true,
		Subcommand:    "sync",
	}

	if a.Run() != 0 {
		t.Errorf("sync subcommand failed with exit code 1, expected 0")
	}
}

func TestNonexistantSubcommand(t *testing.T) {
	a := &App{}

	if a.Run() == 0 {
		t.Errorf("sync subcommand erroneously succeeded with exit code 0, expected 1")
	}
}
