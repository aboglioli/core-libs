package events

import (
	"context"
	"encoding/json"
	"time"

	"github.com/nats-io/nats.go"
)

var _ Publisher = (*NatsEventBus)(nil)
var _ Subscriber = (*NatsEventBus)(nil)

type NatsEvent struct {
	Id        string      `json:"id"`
	EntityId  string      `json:"entity_id"`
	Topic     string      `json:"topic"`
	Payload   interface{} `json:"payload"`
	Timestamp time.Time   `json:"timestamp"`
}

type NatsEventBus struct {
	consumerGroup string
	conn          *nats.Conn
}

func NewNatsEventBus(
	consumerGroup string,
	conn *nats.Conn,
) *NatsEventBus {
	return &NatsEventBus{
		consumerGroup: consumerGroup,
		conn:          conn,
	}
}

func (eb *NatsEventBus) Publish(ctx context.Context, events ...*Event) error {
	for _, event := range events {
		natsEvent := NatsEvent{
			Id:        event.Id(),
			EntityId:  event.EntityId(),
			Topic:     event.Topic(),
			Payload:   event.Payload(),
			Timestamp: event.Timestamp(),
		}

		msg, err := json.Marshal(&natsEvent)
		if err != nil {
			return &ErrSerializingEvent{event, err}
		}

		if err := eb.conn.Publish(event.Topic(), msg); err != nil {
			return &ErrPublishingEvent{event, err}
		}
	}

	return nil
}

func (eb *NatsEventBus) Subscribe(ctx context.Context, subject string, handler Handler) error {
	_, err := eb.conn.QueueSubscribe(subject, eb.consumerGroup, func(msg *nats.Msg) {
		var natsEvent NatsEvent
		if err := json.Unmarshal(msg.Data, &natsEvent); err != nil {
			panic(&ErrDeserializingEvent{err})
		}

		event, err := NewEvent(
			natsEvent.Id,
			natsEvent.EntityId,
			natsEvent.Topic,
			natsEvent.Payload,
			natsEvent.Timestamp,
		)
		if err != nil {
			panic(err)
		}

		handler.Handle(ctx, event)
	})
	if err != nil {
		return &ErrSubscribingToSubject{subject, err}
	}

	return nil
}
