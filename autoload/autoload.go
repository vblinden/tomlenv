package autoload

import "github.com/vblinden/tomlenv"

// init autoload function, import the path to autoload the env.toml file in env
func init() {
	err := tomlenv.Load()
	if err != nil {
		panic(err)
	}
}
