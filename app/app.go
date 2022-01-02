package app

import (
	"fmt"
	"os"

	"github.com/jgrancell/metasync/configuration"
	"github.com/jgrancell/metasync/utils"
)

type App struct {
	Configuration *configuration.Configuration
	Debug         bool
	Dryrun        bool
	ShowDiffs     bool
	Subcommand    string
	Verbose       bool
	Version       string
}

func (a *App) Run() int {
	// Running our subcommand
	if a.Subcommand == "sync" {
		// TODO: add app.Sync() or similar here
		return a.Sync(".")
	} else {
		fmt.Println("Specified subcommand", a.Subcommand, "is not valid.")
		return 1
	}
}

func (a *App) Sync(target string) int {
	a.LogDebug("Beginning repository sync.")

	repo := &Repository{
		Ref:                    a.Configuration.SourceRepoRef,
		RefType:                a.Configuration.SourceRepoRefType,
		RequiresAuthentication: false,
		ShowDiffs:              a.ShowDiffs,
		TemplatesDirectory:     a.Configuration.SourceTemplatePath,
		Url:                    a.Configuration.SourceRepository,
	}

	// Preparation: we setup our starting and target directory variables
	if err := utils.MoveToTarget(target); err != nil {
		return utils.VisualizeError(err)
	}

	// Step 1: we create a temporary directory
	tempdir, err := os.MkdirTemp("", "metasync-git-*")
	defer os.RemoveAll(tempdir)
	if err != nil {
		return utils.VisualizeError(err)
	}

	// Step 2: we clone down the source repository and checkout the target ref
	a.LogDebug("Fetching source templates from remote repository.")
	if err = repo.Clone(tempdir, a.Debug); err != nil {
		return utils.VisualizeError(err)
	}

	a.LogDebug("Checking out request ref " + repo.RefType + "/" + repo.Ref + " for source repository.")

	if err = repo.Checkout(); err != nil {
		return utils.VisualizeError(err)
	}
	a.LogDebug("Setup of source template repository complete.")

	// Step 3: we determine if there are content differences between each template file and local file
	a.LogDebug("Parsing source template files to find changes.")
	if err := repo.FindSyncCandidates(); err != nil {
		return utils.VisualizeError(err)
	}

	// Step 4: we update any sync candidate files if we don't have dryrun enabled
	if len(repo.CandidateFiles) == 0 {
		return 0
	} else {
		if a.Dryrun {
			fmt.Println()
			fmt.Println("*** Dryrun mode is enabled. Exiting without making changes.")
			return 0
		} else {
			fmt.Println()
			a.LogDebug("Updating all out-of-sync files.")
			err := repo.ExecuteCandidateSync()
			if err == nil {
				fmt.Println("Success! All files have been synced.")
			}
			return utils.VisualizeError(err)
		}
	}
}

func (a *App) LogDebug(text string) {
	if a.Debug {
		fmt.Println(text)
	}
}
