package events

import (
	"context"
)

type Handler interface {
	Handle(ctx context.Context, event *Event) error
}

type Subscriber interface {
	Subscribe(ctx context.Context, topic string, handler Handler) error
}
