package app

import (
	"bytes"
	"io/ioutil"
	"os"

	diff "github.com/sergi/go-diff/diffmatchpatch"
)

type SyncCandidate struct {
	Diff           string
	Name           string
	SourceContents []byte
	SourceFile     string
	TargetContents []byte
	TargetFile     string
}

func (c *SyncCandidate) TargetExists() bool {
	_, err := os.Stat(c.TargetFile)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func (c *SyncCandidate) RequiresSync() (bool, error) {
	if c.TargetExists() {
		var err error
		c.SourceContents, err = c.LoadContents(c.SourceFile)
		if err != nil {
			return false, err
		}

		c.TargetContents, err = c.LoadContents(c.TargetFile)
		if err != nil {
			return false, err
		}

		check := diff.New()
		diffs := check.DiffMain(string(c.TargetContents), string(c.SourceContents), false)
		c.Diff = check.DiffPrettyText(diffs)

		if bytes.Equal(c.SourceContents, c.TargetContents) {
			return false, nil
		}
	}
	return true, nil
}

func (c *SyncCandidate) LoadContents(file string) ([]byte, error) {
	contents, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	return contents, nil
}
