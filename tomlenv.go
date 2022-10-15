package tomlenv

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

// Load the default "env.toml" into your environment
func Load() error {
	err := LoadFile("env.toml")

	return err
}

// LoadFile wil load one or multiple files into your environment
func LoadFile(filenames ...string) error {
	if len(filenames) == 0 {
		return fmt.Errorf("You must provide atleast one filename.")
	}

	for _, filename := range filenames {
		config, err := toml.LoadFile(filename)
		if err != nil {
			panic(fmt.Sprintf("Can't read environment variables from %s file.", filename))
		}

		loadInEnvironment(config)
	}

	return nil
}

// LoadString will load environment in based on a given TOML string
func LoadString(content string) error {
	config, err := toml.Load(content)
	if err != nil {
		return err
	}

	loadInEnvironment(config)

	return nil
}

func loadInEnvironment(config *toml.Tree) {
	configMap := config.ToMap()

	for i := range configMap {
		if _, ok := configMap[i].(map[string]interface{}); ok {
			configMapValue := configMap[i].(map[string]interface{})
			for j := range configMapValue {
				_ = os.Setenv(
					fmt.Sprintf("%s.%v", i, j),
					fmt.Sprintf("%v", configMapValue[j]),
				)
			}

			continue
		}

		_ = os.Setenv(i, fmt.Sprintf("%v", configMap[i]))
	}
}
