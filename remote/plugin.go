package remote

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"plugin"

	"github.com/pkg/errors"
	"github.com/adrg/xdg"
	"github.com/google/uuid"
)

// Remote is the implementation of Plugin in order
// to locate and pull remote plugin locally.
// It supports various protocol
type Remote struct {
	// URL is the base URL where the plugin
	// is located
	URL string
}

// Lookup will try to GET the plugin based on the attributes
// of the Remote structure
func (r Remote) Open(path string) (*plugin.Plugin, error) {
	res, err := http.Get(path)
	if err != nil {
		return nil, errors.Wrap(err, "unable to GET the URL")
	}
	body := res.Body
	defer body.Close()
	cacheFilePath, err := xdg.CacheFile(fmt.Sprintf("plugin-%s.so", uuid.New().String()))
	dest, err := os.Create(cacheFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create plugin destination file")
	}
	defer dest.Close()

	if _, err := io.Copy(dest, body); err != nil {
		return nil, errors.Wrap(err, "unable to copy plugin to its destination file")
	}

	return plugin.Open(cacheFilePath)
}
