package golazy

import (
	"sync"
	"time"
)

// withLoader is the internal implementation of Lazy[T] that supports a loader
// function, optional TTL-based per-context caching, and preloaded values.
type withLoader[T any] struct {
	values  map[any]T
	times   map[any]*time.Time
	ttl     time.Duration
	withTTL bool
	loader  LazyFunc[T]
	mu      *sync.Mutex
}

// newWithLoader constructs a withLoader that will call loader on demand. If
// withTTL is true, returned values are cached for the provided ttl duration.
func newWithLoader[T any](loader LazyFunc[T], withTTL bool, ttl time.Duration) *withLoader[T] {
	return &withLoader[T]{
		values:  make(map[any]T),
		times:   make(map[any]*time.Time),
		loader:  loader,
		ttl:     ttl,
		withTTL: withTTL,
		mu:      &sync.Mutex{},
	}
}

// newWithLoaderPreloaded constructs a withLoader pre-populated with a value
// for the provided ctx. This is useful when you already have a value and want
// to expose it through the Lazy API.
func newWithLoaderPreloaded[T any](loader LazyFunc[T], value T, ctx any, withTTL bool, ttl time.Duration) *withLoader[T] {
	values := make(map[any]T)
	values[ctx] = value

	return &withLoader[T]{
		values:  values,
		times:   make(map[any]*time.Time),
		loader:  loader,
		ttl:     ttl,
		withTTL: withTTL,
		mu:      &sync.Mutex{},
	}
}

// Value returns the cached value for ctx if present and not expired. Otherwise
// it calls the loader (if set) to obtain the value, caches it, and returns it.
// The method is safe for concurrent use because it synchronizes using a mutex.
func (l *withLoader[T]) Value(ctx any) (T, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	needsRefresh := l.withTTL && l.times[ctx] != nil && time.Since(*l.times[ctx]) > l.ttl

	if v, ok := l.values[ctx]; ok && !needsRefresh {
		return v, nil
	}

	if l.loader == nil {
		return l.values[nil], nil
	}

	v, err := l.loader(ctx)
	if err == nil {
		l.values[ctx] = v
	}

	if l.withTTL {
		t := time.Now()
		l.times[ctx] = &t
	}

	return v, err
}

// Clear removes the cached value for the given context.
func (l *withLoader[T]) Clear(ctx any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.values, ctx)
	delete(l.times, ctx)
}

// ClearAll clears the entire cache maintained by the withLoader.
func (l *withLoader[T]) ClearAll() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.values = make(map[any]T)
	l.times = map[any]*time.Time{}
}
