package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStyId(t *testing.T) {
	t.Run("new id", func(t *testing.T) {
		// Valid
		id1, err := NewStrId("id#01")
		assert.NoError(t, err)
		assert.Equal(t, "id#01", id1.Value())

		id2, err := NewStrId("id#02")
		assert.NoError(t, err)
		assert.Equal(t, "id#02", id2.Value())

		assert.True(t, id1.Equals(id1))
		assert.False(t, id1.Equals(id2))

		// Invalid
		_, err = NewStrId("")
		assert.Error(t, err)
	})

	t.Run("generate uuid", func(t *testing.T) {
		id := GenerateUuid()
		assert.Greater(t, len(id.Value()), 10)
	})
}
