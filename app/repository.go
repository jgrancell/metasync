package app

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"

	git "github.com/go-git/go-git/v5"
	plumbing "github.com/go-git/go-git/v5/plumbing"

	"github.com/jgrancell/metasync/utils"
)

type Repository struct {
	CandidateFiles         []*SyncCandidate
	FilesystemPath         string
	RawRepository          *git.Repository
	Ref                    string
	RefType                string
	RequiresAuthentication bool
	ShowDiffs              bool
	TemplatesDirectory     string
	Url                    string
}

func (r *Repository) Clone(tempdir string, verbose bool) error {
	repo, err := git.PlainClone(tempdir, false, &git.CloneOptions{
		URL: r.Url,
	})
	r.RawRepository = repo
	r.FilesystemPath = tempdir
	return utils.ReturnError("failed to clone source repository", err)
}

func (r *Repository) Checkout() error {
	w, err := r.RawRepository.Worktree()
	if err != nil {
		return utils.ReturnError("unable to open source repository worktree", err)
	}

	var opts *git.CheckoutOptions
	if r.RefType == "branch" {
		opts = &git.CheckoutOptions{
			Branch: plumbing.ReferenceName("refs/remotes/origin/" + r.Ref),
		}
	} else {
		opts = &git.CheckoutOptions{
			Hash: plumbing.NewHash(r.Ref),
		}
	}
	err = w.Checkout(opts)
	return utils.ReturnError("unable to checkout target ref", err)
}

func (r *Repository) FindSyncCandidates() error {
	sourceFiles, err := r.ListFiles()
	if err != nil {
		return utils.ReturnError("unable to list files from source repository", err)
	}

	cwd, err := os.Getwd()
	if err != nil {
		return utils.ReturnError("unable to determine the current working directory", err)
	}

	for _, source := range sourceFiles {

		c := &SyncCandidate{
			Name:       source.Name(),
			SourceFile: filepath.Join(r.FilesystemPath, r.TemplatesDirectory, source.Name()),
			TargetFile: filepath.Join(cwd, source.Name()),
		}

		ok, err := c.RequiresSync()
		if err != nil {
			fmt.Println("Sync candidate", c.Name, "could not be read properly. Skipping.")
		}
		if ok {
			r.CandidateFiles = append(r.CandidateFiles, c)
		}
	}

	r.PrintCandidates()
	return nil
}

func (r *Repository) ExecuteCandidateSync() error {
	for _, c := range r.CandidateFiles {
		err := c.Sync()
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Repository) ListFiles() ([]fs.FileInfo, error) {
	files, err := ioutil.ReadDir(r.FilesystemPath + "/" + r.TemplatesDirectory)
	if err != nil {
		return nil, err
	}
	return files, nil
}

func (r *Repository) PrintCandidates() {
	if len(r.CandidateFiles) == 0 {
		fmt.Println("All files are in sync.")
		return
	}

	if !r.ShowDiffs {
		fmt.Println("Out of sync files:")
	}

	for _, c := range r.CandidateFiles {
		if r.ShowDiffs {
			fmt.Println()
			fmt.Println(c.Name + ":")
			if len(c.Diff) > 0 {
				fmt.Println(c.Diff)
			} else {
				fmt.Println("File does not exist in target.")
			}
			fmt.Println("")
			fmt.Println("----------------------------------------")
		} else {
			fmt.Println("  -", c.Name)
		}
	}
}
