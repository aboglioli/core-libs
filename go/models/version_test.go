package models

import (
	"testing"

	"github.com/aboglioli/core-libs/go/types"
	"github.com/stretchr/testify/assert"
)

func TestVersion(t *testing.T) {
	t.Run("init version", func(t *testing.T) {
		v1 := InitVersion()
		assert.Equal(t, int64(1), v1.Value())

		v2 := v1.Incr()
		assert.Equal(t, int64(1), v2.Value())

		assert.True(t, v1.Equals(v2))
	})

	t.Run("existing version", func(t *testing.T) {
		assert.Error(t, types.UnwrapError(NewVersion(0)))

		v1, err := NewVersion(3)
		assert.NoError(t, err)

		assert.Equal(t, int64(3), v1.Value())

		v2 := v1.Incr()
		assert.Equal(t, int64(4), v2.Value())

		v3 := v2.Incr()
		assert.Equal(t, int64(4), v3.Value())

		assert.False(t, v1.Equals(v2))
		assert.True(t, v2.Equals(v3))
	})
}
