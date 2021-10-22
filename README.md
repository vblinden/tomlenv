## TOMLenv
This simple package will read all the variables from your `.toml` file and put them in the environment so you can easily access them via `os.Getenv("varname")`.

### Installation
```shell
go get -u github.com/vblinden/tomlenv
```

### Example
env.toml
```toml
token = "abcdef123456"

[database]
username = "john"
password = "secret"
```

main.go
```go
package main

import "github.com/vblinden/tomlenv"

func main() {
    // To load a file with name "env.toml"...
    tomlenv.Load()

    // To load a file with other names...
    tomlenv.LoadFile("custom.toml")

    // To load a string...
    tomlenv.LoadString("token = \"supersecret\"")

    // Will print "abcdef123456"...
    fmt.Printf(os.Getenv("token"))
    
    // To load a string...
    tomlenv.LoadString("[database]  username = \"john\"")

    // Will print "john"...
    fmt.Printf(os.Getenv("database.username"))
}
```
