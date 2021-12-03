package configuration

import (
	"fmt"
	"io/ioutil"

	"github.com/jgrancell/metasync/utils"
	yaml "gopkg.in/yaml.v2"
)

type Configuration struct {
	CloneViaSSH       bool
	ConfigurationPath string
	TargetRef         string
	TargetRefType     string
	TargetRepository  string
	TargetRootPath    string
}

func (c *Configuration) Load() error {
	if err := utils.CheckFileExists(c.ConfigurationPath); err != nil {
		return fmt.Errorf("the specified configuration file %s does not exist", c.ConfigurationPath)
	}

	return c.FetchConfig()
}

func (c *Configuration) FetchConfig() error {
	contents, err := ioutil.ReadFile(c.ConfigurationPath)
	if err != nil {
		return err
	}
	if err := yaml.Unmarshal(contents, c); err != nil {
		return err
	}
	return nil
}
