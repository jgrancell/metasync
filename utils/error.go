package utils

import "fmt"

func ReturnError(text string, err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%s: %s", text, err.Error())
}

func VisualizeError(err error) int {
	if err != nil {
		fmt.Println(err.Error())
		return 1
	}
	return 0
}
