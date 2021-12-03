package app

import "testing"

func TestSyncSubcommand(t *testing.T) {
	a := &App{
		Subcommand: "sync",
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
