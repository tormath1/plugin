# Plugin

[![PkgGoDev](https://pkg.go.dev/badge/github.com/tormath1/plugin)](https://pkg.go.dev/github.com/tormath1/plugin)
![Go](https://github.com/tormath1/plugin/workflows/Go/badge.svg)

The goal of this lib is to provide a set of API to load plugins from different locations (remote, local, etc.)

## Example

Let's suppose one plugin is stored locally:

```go
package main

import (
        "fmt"

        "github.com/tormath1/plugin"
        "github.com/tormath1/plugin/factory"
)

func main() {
        loader, _ := factory.Get(plugin.Local)
        p, _ := loader.Open("/tmp/plugin.so")
        fmt.Printf("%T\n", p)
}
```

Then:
```shell
$ go run main.go
*plugin.Plugin
```

## Supported plugin types

- `local`: basically, the default behavior of plugins
- `remote`: fetch a plugin from a remote location and download it in the `$XDG_CACHE` directory
- `embedded`: extract a plugin packaged with `pkger` from a Go binary 

## Contributing

Just create a new dir in the `plugin/` directory:

```shell
mkdir my-new-plugin-type
```

Then, you just need to define a structure implementing the `plugin.Plugin` interface:

```go
package local

import "plugin"

// Local is the implementation of Plugin in order
// to lookup for local plugin
type Local struct{}

// Open will try to find the plugin locally, it's the default
// behavior
func (l Local) Open(path string) (*plugin.Plugin, error) {
        return plugin.Open(path)
}
```

Finally, it's required to add the `my-new-plugin-type` to the list of constants in `plugin.Type`, generate the new content: `make` and finally add the mapping in the `plugin/factory.go`.
