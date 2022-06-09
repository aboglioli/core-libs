package events

import (
	"context"
	"strings"
	"sync"
)

var _ Publisher = (*LocalEventBus)(nil)
var _ Subscriber = (*LocalEventBus)(nil)

type subscription struct {
	subject string
	handler Handler
}

type LocalEventBus struct {
	mux           sync.Mutex
	subscriptions []*subscription
}

func NewLocalEventBus() *LocalEventBus {
	return &LocalEventBus{
		subscriptions: make([]*subscription, 0, 10),
	}
}

func subjectHasTopic(subject string, topic string) bool {
	if subject == topic {
		return true
	}

	subjectParts := strings.Split(subject, ".")
	topicParts := strings.Split(topic, ".")

	if len(subjectParts) != len(topicParts) {
		return false
	}

	for i, subjectPart := range subjectParts {
		topicPart := topicParts[i]

		if subjectPart != "*" && subjectPart != topicPart {
			return false
		}
	}

	return true
}

func (eb *LocalEventBus) Publish(ctx context.Context, events ...*Event) error {
	for _, event := range events {
		for _, subscription := range eb.subscriptions {
			if subjectHasTopic(subscription.subject, event.Topic()) {
				if err := subscription.handler.Handle(ctx, event); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (eb *LocalEventBus) Subscribe(ctx context.Context, subject string, handler Handler) error {
	eb.mux.Lock()
	defer eb.mux.Unlock()

	eb.subscriptions = append(eb.subscriptions, &subscription{
		subject: subject,
		handler: handler,
	})

	return nil
}
