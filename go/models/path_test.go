package models

import (
	"testing"

	"github.com/aboglioli/libs/go/types"
	"github.com/stretchr/testify/assert"
)

func TestPath(t *testing.T) {
	t.Run("new path", func(t *testing.T) {
		// Simple
		path, err := NewPath("my.path", ".")
		assert.NoError(t, err)
		assert.Equal(t, []string{"my", "path"}, path.Parts())
		assert.Equal(t, ".", path.Separator())
		assert.Len(t, path.Wildcards(), 0)
		assert.Equal(t, "my.path", path.String())

		// With wildcards
		path, err = NewPath("my.path", ".", "*", ">")
		assert.NoError(t, err)
		assert.Equal(t, []string{"my", "path"}, path.Parts())
		assert.Equal(t, ".", path.Separator())
		assert.Len(t, path.Wildcards(), 2)
		assert.Equal(t, "my.path", path.String())

		// Single part
		path, err = NewPath("my_path", "#", "*", ">")
		assert.NoError(t, err)
		assert.Equal(t, []string{"my_path"}, path.Parts())
		assert.Equal(t, "#", path.Separator())
		assert.Len(t, path.Wildcards(), 2)
		assert.Equal(t, "my_path", path.String())

		// Invalid character
		_, err = NewPath("my.p@th", ".")
		assert.Error(t, err)
	})

	t.Run("equals", func(t *testing.T) {
		path1 := types.Unwrap(NewPath("one.two.three", "."))
		path2 := types.Unwrap(NewPath("one.three.two", "."))
		path3 := types.Unwrap(NewPath("two.one.three", "."))

		assert.True(t, path1.Equals(path1))
		assert.False(t, path1.Equals(path2))
		assert.False(t, path2.Equals(path3))
	})
}
