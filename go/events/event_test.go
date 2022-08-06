package events

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEvent(t *testing.T) {
	t.Run("create", func(t *testing.T) {
		type data struct {
			Num int8 `json:"num"`
		}

		event, err := CreateEvent(
			"entity#01",
			"topic",
			&data{123},
		)

		assert.NoError(t, err)

		assert.Equal(t, "entity#01", event.EntityId())
		assert.Equal(t, "topic", event.Topic())
	})

	t.Run("unmarshal struct", func(t *testing.T) {
		type data struct {
			Num int8 `json:"num"`
		}

		event, _ := CreateEvent(
			"entity#01",
			"topic",
			&data{123},
		)

		var p data
		err := event.UnmarshalPayload(&p)
		assert.NoError(t, err)

		assert.Equal(t, data{123}, p)
	})

	t.Run("unmarshal json", func(t *testing.T) {
		jsonBytes, _ := json.Marshal(map[string]int64{
			"num": 123,
		})

		event, _ := CreateEvent(
			"entity#01",
			"topic",
			jsonBytes,
		)

		type data struct {
			Num int8 `json:"num"`
		}

		var p data
		err := event.UnmarshalPayload(&p)
		assert.NoError(t, err)

		assert.Equal(t, data{123}, p)
	})
}
