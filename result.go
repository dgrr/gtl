package gtl

import (
	"fmt"
)

// Result represents the result of a function.
//
// A result has an expected value, which is T
// and an unexpected value, which is E. E is often used as error.
type Result[T any] struct {
	value T
	err   error
}

// MakeResult returns a Result from 2 values. It is useful when we want
// to capture old-style function returning.
func MakeResult[T any](v T, err error) (r Result[T]) {
	return r.Any(v, err)
}

// Ok returns a Result holding an expected value.
//
// If `r` was previously holding an unexpected value that is discarded.
func (r Result[T]) Ok(v T) Result[T] {
	r.value = v
	r.err = nil

	return r
}

// IsOk returns whether the Result has the unexpected value set or not.
func (r Result[T]) IsOk() bool {
	return r.err == nil
}

// Err returns a Result holding an unexpected value (an error).
func (r Result[T]) Err(e error) Result[T] {
	r.err = e

	return r
}

// Any takes from input either T or E. If E is defined, V is not used.
func (r Result[T]) Any(value T, err error) Result[T] {
	if err == nil {
		return r.Ok(value)
	}

	return r.Err(err)
}

// Or sets the value `value` only if the error was previously set.
// That means: In case the Result has failed, set this alternative value.
//
// If the Result is holding an error, the error is cleared out and `value` is set.
func (r Result[T]) Or(v T) Result[T] {
	if r.err != nil {
		return r.Ok(v)
	}

	return r
}

// Map passes the holding value to `fn` if no error is being held.
func (r Result[T]) Map(fn func(T) T) Result[T] {
	if r.err == nil {
		r.value = fn(r.value)
	}

	return r
}

// Unwrap unwraps the Result returning the value, if any.
func (r Result[T]) Unwrap() T {
	return r.Get()
}

// Expect returns the holding value. If `r` is holding an error, the function panics
// with a message of the format: `str: error`.
func (r Result[T]) Expect(str string) T {
	if r.err != nil {
		panic(fmt.Sprintf("%s: %s", str, r.err))
	}

	return r.value
}

// Get returns the expected value.
func (r Result[T]) Get() T {
	return r.value
}

// Err returns the unexpected value.
func (r Result[T]) Error() error {
	return r.err
}

// Both returns the value and the error.
func (r Result[T]) Both() (T, error) {
	return r.Get(), r.Error()
}

// Then is executed if Result is not holding an unexpected error.
func (r Result[T]) Then(fn func(T)) Result[T] {
	if r.err == nil {
		fn(r.value)
	}

	return r
}

// ThenE is executed if Result is not holding an unexpected error.
//
// The function must return en error. The error can be nil.
// The returned error will be set to the Result's error.
func (r Result[T]) ThenE(fn func(T) error) Result[T] {
	if r.err == nil {
		return r.Err(fn(r.value))
	}

	return r
}

// Else is executed if Result is holding an error.
func (r Result[T]) Else(fn func(error)) Result[T] {
	if r.err != nil {
		fn(r.err)
	}

	return r
}

// ElseE is executed if Result is holding an error.
//
// The function must return an error. The error can be nil.
// The returned error will be set to the Result's error.
func (r Result[T]) ElseE(fn func(error) error) Result[T] {
	if r.err != nil {
		return r.Err(fn(r.err))
	}

	return r
}
