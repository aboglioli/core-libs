package cache

import (
	"context"

	"github.com/aboglioli/libs/go/errors"
)

var (
	ErrCacheInternal = errors.Define("cache.internal")
)

type Cache[K any, V any] interface {
	Get(ctx context.Context, k K) (V, error)
	Set(ctx context.Context, k K, v V) error
	Delete(ctx context.Context, k K) error
}
