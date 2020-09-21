package factory

import (
	"fmt"

	"github.com/tormath1/plugin"
	"github.com/tormath1/plugin/embedded"
	"github.com/tormath1/plugin/local"
	"github.com/tormath1/plugin/remote"
)

var (
	plugins = map[plugin.Type]plugin.Plugin{
		plugin.Local:    local.Local{},
		plugin.Remote:   remote.Remote{},
		plugin.Embedded: embedded.Embedded{},
	}
)

// Get returns the specific Plugin for t
func Get(t string) (plugin.Plugin, error) {
	ts, err := plugin.TypeString(t)
	if err != nil {
		return nil, fmt.Errorf("no plugin type for: %s", ts)
	}

	p, ok := plugins[ts]
	if !ok {
		return nil, fmt.Errorf("no plugin defined for: %s", ts)
	}
	return p, nil
}
