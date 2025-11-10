package golazy

// static is a simple Lazy[T] implementation that always returns a fixed value.
// It satisfies the Lazy[T] interface but does not perform any loading or
// caching behavior (Clear and ClearAll are no-ops).
type static[T any] struct {
	value T
}

// Value returns the static value and a nil error.
func (l static[T]) Value(ctx any) (T, error) {
	return l.value, nil
}

// Clear is a no-op for static values.
func (l static[T]) Clear(ctx any) {
	// nothing
}

// ClearAll is a no-op for static values.
func (l static[T]) ClearAll() {
	// nothing
}
