## TOMLenv
This simpel package will read all the variables from your `.toml` file and put them in the environment so you can easily access them via `os.Getenv("varname")`.


### Installation
```shell
go get -u github.com/vblinden/tomlenv
```

### Example
env.toml
```toml
token = "supersecret"
```

main.go
```go
package main

import "github.com/vblinden/tomlenv"

func main() {
    // To load the default env.toml file...
    tomlenv.Load()

    // To load a custom .toml file...
    tomlenv.Load("custom.toml")

    // Will print "supersecret"...
    fmt.Printf(os.Getenv("token"))
}
```