package local

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLocal(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		l := Local{}
		p, err := l.Open("../testdata/plugin.so")
		require.Nil(t, err)
		f, err := p.Lookup("Function")
		require.Nil(t, err)
		assert.Equal(t, f.(func() int)(), 5)
	})
	t.Run("Error", func(t *testing.T) {
		l := Local{}
		p, err := l.Open("../testdata/p")
		require.Nil(t, p)
		assert.Error(t, err)
	})

}
