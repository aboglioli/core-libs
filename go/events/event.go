package events

import (
	"encoding/json"
	"reflect"
	"time"

	"github.com/google/uuid"
)

type Event struct {
	id        string
	entityId  string
	topic     string
	payload   interface{}
	timestamp time.Time
}

func NewEvent(
	id string,
	entityId string,
	topic string,
	payload interface{},
	timestamp time.Time,
) (*Event, error) {
	if len(id) == 0 {
		return nil, ErrInvalidEvent
	}

	if len(entityId) == 0 {
		return nil, ErrInvalidEvent
	}

	if len(topic) == 0 {
		return nil, ErrInvalidEvent
	}

	if payload == nil {
		return nil, ErrInvalidEvent
	}

	return &Event{
		id,
		entityId,
		topic,
		payload,
		timestamp,
	}, nil
}

func CreateEvent(
	entityId string,
	topic string,
	payload interface{},
) (*Event, error) {
	return NewEvent(
		uuid.New().String(),
		entityId,
		topic,
		payload,
		time.Now(),
	)
}

func (e *Event) Id() string {
	return e.id
}

func (e *Event) EntityId() string {
	return e.entityId
}

func (e *Event) Topic() string {
	return e.topic
}

func (e *Event) Payload() interface{} {
	return e.payload
}

func (e *Event) UnmarshalPayload(v interface{}) error {
	if reflect.TypeOf(e.payload) == reflect.TypeOf(v) {
		reflect.ValueOf(v).Elem().Set(reflect.ValueOf(e.payload).Elem())

		return nil
	}

	if payload, ok := e.payload.(json.RawMessage); ok {
		if err := json.Unmarshal(payload, v); err != nil {
			return &ErrDeserializingPayload{e, err}
		}

		return nil
	}

	if payload, ok := e.payload.([]byte); ok {
		if err := json.Unmarshal(payload, v); err != nil {
			return &ErrDeserializingPayload{e, err}
		}

		return nil
	}

	return ErrUnknownPayload
}

func (e *Event) Timestamp() time.Time {
	return e.timestamp
}
