package events

import (
	"context"
)

type Publisher interface {
	Publish(ctx context.Context, events ...*Event) error
}
