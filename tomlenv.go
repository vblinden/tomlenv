package tomlenv

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml"
)

func Load(filenames ...string) {
	if len(filenames) == 0 {
		filenames = []string{"env.toml"}
	}

	for _, filename := range filenames {
		config, err := toml.LoadFile(filename)
		if err != nil {
			panic(fmt.Sprintf("Can't read env variables from %s file.", filename))
		}

		for _, key := range config.Keys() {
			os.Setenv(key, fmt.Sprintf("%v", config.GetDefault(key, "")))
		}
	}
}
