package configuration

import (
	"os"
	"testing"
)

func TestFetchConfig(t *testing.T) {
	pwd, _ := os.Getwd()

	os.Chdir("../testdata")
	c := Configuration{
		ConfigurationPath: ".metasync.yml",
	}
	c.Load()
	os.Chdir(pwd)
}
