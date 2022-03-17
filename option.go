package result

type Option[T any] struct {
	value *T
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: &value}
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func (opt Option[T]) Some() bool {
	return opt.value != nil
}

func (opt Option[T]) Value() T {
	return *opt.value
}

func (opt Option[T]) Unwrap() T {
	if opt.value == nil {
		panic("unwrap on nil value")
	}
	return *opt.value
}

func (opt Option[T]) None() bool {
	return opt.value == nil
}

func (opt Option[T]) BeSome(value T) Option[T] {
	opt.value = &value
	return opt
}

func (opt Option[T]) BeNone() Option[T] {
	opt.value = nil
	return opt
}
