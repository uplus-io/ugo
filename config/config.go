package config

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"os"
)

func LoadYaml(params []string, config interface{}) error {
	if len(params) < 2 {
		return errors.New("must give config file path")
	}

	configPath := params[1]

	switch config.(type) {
	case Configurable:
		config.(Configurable).SetConfigPath(configPath)
	default:
		return errors.New("config entity must implement ugo.config.Configurable")
	}

	fmt.Printf("ready read config file[%s]\n", configPath)

	file, err := os.OpenFile(configPath, os.O_RDONLY, 0600)
	if err != nil {
		return err
	}
	yaml.NewDecoder(file).Decode(config)
	return nil
}
