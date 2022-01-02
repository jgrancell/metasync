package utils

import (
	"os"
)

func CheckFileExists(path string) error {
	if _, err := os.Stat(path); err != nil {
		return err
	}
	return nil
}
