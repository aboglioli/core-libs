package events

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"time"

	"github.com/aboglioli/libs/go/collections"
	"github.com/aboglioli/libs/go/errors"
)

var (
	ErrInvalidEvent              = errors.Define("event.invalid")
	ErrEventPayloadSerialization = errors.Define("event.payload_serialization")
	ErrEventInternal             = errors.Define("event.internal")
)

// Publisher and subscriber
type Publisher interface {
	Publish(ctx context.Context, events ...*Event) error
}

type Handler interface {
	Handle(ctx context.Context, event *Event) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, subject string, handler Handler) error
}

// Event
type Event struct {
	id        string
	entityId  string
	topic     string
	payload   []byte
	timestamp time.Time
}

func NewEvent(
	id string,
	entityId string,
	topic string,
	payload []byte,
	timestamp time.Time,
) (*Event, error) {
	m := collections.WithMetadata("id", id).
		And("entity_id", entityId).
		And("topic", topic).
		And("payload", payload).
		And("timestamp", timestamp)

	if len(id) == 0 {
		return nil, errors.New(ErrInvalidEvent, "payload id is empty", m)
	}

	if len(entityId) == 0 {
		return nil, errors.New(ErrInvalidEvent, "payload entity_id is empty", m)
	}

	if len(topic) == 0 {
		return nil, errors.New(ErrInvalidEvent, "payload topic is empty", m)
	}

	if len(payload) == 0 {
		return nil, errors.New(ErrInvalidEvent, "payload is empty", m)
	}

	return &Event{
		id:        id,
		entityId:  entityId,
		topic:     topic,
		payload:   payload,
		timestamp: timestamp,
	}, nil
}

func CreateEvent(
	entityId string,
	topic string,
	payload interface{},
) (*Event, error) {
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(
			ErrEventPayloadSerialization,
			err,
			"could not serialize event payload",
			collections.WithMetadata("entity_id", entityId).
				And("topic", topic).
				And("payload", payload),
		)
	}

	return NewEvent(
		uuid.New().String(),
		entityId,
		topic,
		payloadBytes,
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

func (e *Event) Payload() []byte {
	return e.payload
}

func (e *Event) UnmarshalPayload(v interface{}) error {
	if err := json.Unmarshal(e.payload, v); err != nil {
		return errors.Wrap(
			ErrEventPayloadSerialization,
			err,
			"could not deserialize event payload",
		)
	}

	return nil
}

func (e *Event) Timestamp() time.Time {
	return e.timestamp
}
