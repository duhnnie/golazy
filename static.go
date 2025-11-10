package golazy

type static[T any] struct {
	value T
}

func (l static[T]) Value(ctx any) (T, error) {
	return l.value, nil
}

func (l static[T]) Reset(ctx any) {
	// nothing
}

func (l static[T]) ResetAll() {
	// nothing
}
