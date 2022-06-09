package events

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	type data struct {
		Msg string
	}

	t.Run("create", func(t *testing.T) {
		event, err := CreateEvent(
			"entity#01",
			"topic.code",
			data{Msg: "Hello World"},
		)
		assert.NoError(t, err)

		assert.Equal(t, "entity#01", event.EntityId())
		assert.Equal(t, "topic.code", event.Topic())
		assert.Greater(t, len(event.Payload()), 0)
	})

	t.Run("payload serialization and deserialization", func(t *testing.T) {
		event, err := CreateEvent(
			"entity#01",
			"topic.code",
			data{Msg: "Hello World"},
		)
		assert.NoError(t, err)

		var d data
		err = event.UnmarshalPayload(&d)
		assert.NoError(t, err)

		assert.Equal(t, data{Msg: "Hello World"}, d)
	})
}
