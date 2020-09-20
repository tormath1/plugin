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
