package utils

import (
	"fmt"
	"os"
)

func MoveToTarget(target string) error {
	cwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("unable to determine current working directory: %s", err.Error())
	}

	if len(target) == 0 || target == "." || target == cwd {
		return nil
	}

	_, err = os.Stat(target)
	if err != nil {
		return fmt.Errorf("unable to find target directory %s: %s", target, err.Error())
	}

	if err = os.Chdir(target); err != nil {
		return fmt.Errorf("unable to move to target %s: %s", target, err.Error())
	}
	return nil
}
