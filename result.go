package result

type Result[T any] struct {
	value *T
	err   error
}

func Ok[T any](value T) Result[T] {
	return Result[T]{
		value: &value,
		err:   nil,
	}
}

func Err[T any](err error) Result[T] {
	return Result[T]{
		value: nil,
		err:   err,
	}
}

func (result Result[T]) Ok() bool {
	return result.err == nil
}

func (result Result[T]) IsErr() bool {
	return result.err != nil
}

func (result Result[T]) Value() T {
	return *result.value
}

func (result Result[T]) Err() error {
	return result.err
}

func (result Result[T]) BeValue(value T) Result[T] {
	result.value = &value
	result.err = nil
	return result
}

func (result Result[T]) BeErr(err error) Result[T] {
	result.value = nil
	result.err = err
	return result
}

// Implement Error interface
func (result Result[T]) Error() string {
	return result.err.Error()
}

type Empty = *struct{}

type EResult = Result[Empty]

func OkResult() Result[Empty] {
	return Ok[Empty](nil)
}

func ErrResult(err error) Result[Empty] {
	return Err[Empty](err)
}
