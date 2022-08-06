package cache

import (
	"context"
	"errors"
)

var (
	ErrCacheInternal = errors.New("internal cache error")
)

type Cache[K any, V any] interface {
	Get(ctx context.Context, k K) (V, error)
	Set(ctx context.Context, k K, v V) error
	Delete(ctx context.Context, k K) error
}
