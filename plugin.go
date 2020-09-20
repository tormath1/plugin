package plugin

import "plugin"

// Plugin is the interface to implement
// for each new types
type Plugin interface {
	Open(string) (*plugin.Plugin, error)
}
