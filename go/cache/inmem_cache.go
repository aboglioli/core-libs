package cache

import (
	"context"
	"sync"
)

var _ Cache[string, interface{}] = (*InMemCache[string, interface{}])(nil)

type InMemCache[K comparable, V any] struct {
	mux   sync.RWMutex
	items map[K]V
}

func NewInMemCache[K comparable, V any]() *InMemCache[K, V] {
	return &InMemCache[K, V]{
		items: make(map[K]V),
	}
}

func (c *InMemCache[K, V]) All() map[K]V {
	c.mux.RLock()
	defer c.mux.RUnlock()

	return c.items
}

func (c *InMemCache[K, V]) Get(ctx context.Context, k K) (V, error) {
	c.mux.RLock()
	defer c.mux.RUnlock()

	if v, ok := c.items[k]; ok {
		return v, nil
	}

	var v V
	return v, nil
}

func (c *InMemCache[K, V]) Set(ctx context.Context, k K, v V) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	c.items[k] = v

	return nil
}

func (c *InMemCache[K, V]) Delete(ctx context.Context, k K) error {
	c.mux.Lock()
	defer c.mux.Unlock()

	delete(c.items, k)

	return nil
}
