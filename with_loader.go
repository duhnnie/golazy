package golazy

import (
	"sync"
	"time"
)

type withLoader[T any] struct {
	values  map[any]T
	times   map[any]*time.Time
	ttl     time.Duration
	withTTL bool
	loader  LazyFunc[T]
	mu      *sync.Mutex
}

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

func (l withLoader[T]) Value(ctx any) (T, error) {
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

func (l withLoader[T]) Reset(ctx any) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.values, ctx)
}

func (l withLoader[T]) ResetAll() {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.values = make(map[any]T)
}
