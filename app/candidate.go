package app

import (
	"bytes"
	"io/fs"
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
	TargetMode     fs.FileMode
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

func (c *SyncCandidate) GetWriteMode(path string) error {
	file, err := os.Stat(path)
	if err != nil {
		return err
	}

	c.TargetMode = file.Mode()
	return nil
}

func (c *SyncCandidate) RequiresSync() (bool, error) {
	var err error
	if c.SourceContents, err = c.LoadContents(c.SourceFile); err != nil {
		return false, err
	}

	if c.TargetExists() {
		// Getting the file mode for the target file
		if err := c.GetWriteMode(c.TargetFile); err != nil {
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
	if err := c.GetWriteMode(c.SourceFile); err != nil {
		return false, err
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

func (c *SyncCandidate) Sync() error {
	return os.WriteFile(c.TargetFile, c.SourceContents, c.TargetMode)
}
