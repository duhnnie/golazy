package golazy

import (
	"time"
)

type Lazy[T any] interface {
	Value(ctx any) (T, error)
	Reset(ctx any)
	ResetAll()
}

type LazyFunc[T any] func(ctx any) (T, error)

func WithLoader[T any](loader LazyFunc[T]) Lazy[T] {
	return newWithLoader(loader, false, 0)
}

func WithLoaderTTL[T any](loader LazyFunc[T], ttl time.Duration) Lazy[T] {
	return newWithLoader(loader, true, ttl)
}

func Preloaded[T any](loader LazyFunc[T], value T, ctx any) Lazy[T] {
	return newWithLoaderPreloaded(loader, value, ctx, false, 0)
}

func PreloadedTTL[T any](loader LazyFunc[T], value T, ctx any, ttl time.Duration) *withLoader[T] {
	return newWithLoaderPreloaded(loader, value, ctx, true, ttl)
}

func NewStatic[T any](value T) *static[T] {
	return &static[T]{
		value: value,
	}
}
