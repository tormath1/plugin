package embedded

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"plugin"

	"github.com/adrg/xdg"
	"github.com/google/uuid"
	"github.com/markbates/pkger"
	"github.com/pkg/errors"
)

type Embedded struct{}

func (e Embedded) Open(path string) (*plugin.Plugin, error) {
	fullPath := fmt.Sprintf("/plugins/%s", path)
	if _, err := pkger.Stat(fullPath); err != nil {
		return nil, errors.Wrap(err, "unable to find plugin in path")
	}

	f, err := pkger.Open(fullPath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to open plugin")
	}
	defer f.Close()

	payload, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, errors.Wrap(err, "unable to extract plugin content")
	}
	cacheFilePath, err := xdg.CacheFile(fmt.Sprintf("plugin-%s.so", uuid.New().String()))
	dest, err := os.Create(cacheFilePath)
	if err != nil {
		return nil, errors.Wrap(err, "unable to create plugin destination file")
	}
	defer dest.Close()

	if _, err := io.Copy(dest, bytes.NewBuffer(payload)); err != nil {
		return nil, errors.Wrap(err, "unable to copy plugin to its destination file")
	}

	return plugin.Open(cacheFilePath)
}
