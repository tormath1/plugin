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
func Get(t plugin.Type) (plugin.Plugin, error) {
	p, ok := plugins[t]
	if !ok {
		return nil, fmt.Errorf("no plugin defined for: %s", t.String())
	}
	return p, nil
}
