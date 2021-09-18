package gtl

import (
	"fmt"
)

// Result represents the result of a function.
//
// A result has an expected value, which is T
// and an unexpected value, which is E. E is often used as error.
type Result[T any] struct {
	v T
	e error
}

// NewResult returns a Result from 2 values. It is useful when we want
// to capture old-style function returning.
func NewResult[T any](v T, err error) (r Result[T]) {
	r.v = v
	if err != nil {
		r.e = err
	}

	return r
}

// Ok returns a Result holding an expected value.
//
// If `r` was previously holding an unexpected value that is discarded.
func (r Result[T]) Ok(v T) Result[T] {
	r.v = v

	return r
}

// IsOk returns whether the Result has the unexpected value set or not.
func (r Result[T]) IsOk() bool {
	return r.e == nil
}

// Err returns a Result holding an unexpected value (an error).
func (r Result[T]) Err(e error) Result[T] {
	r.e = e

	return r
}

// Any takes from input either T or E. If E is defined, V is not used.
func (r Result[T]) Any(v T, e error) Result[T] {
	r.e = e
	if e == nil {
		r.v = v
	}

	return r
}

// Or sets the value `v` only if the error was previously set.
// That means: In case the Result has failed, set this alternative value.
//
// If the Result is holding an error, the error is cleared out and `v` is set.
func (r Result[T]) Or(v T) Result[T] {
	if r.e == nil {
		r = r.Ok(v)
	}

	return r
}

// Map passes the holding value to `fn` if no error is being held.
func (r Result[T]) Map(fn func(T) T) Result[T] {
	if r.e == nil {
		r.v = fn(r.v)
	}

	return r
}

// Unwrap unwraps the Result returning the value, if any.
func (r Result[T]) Unwrap() T {
	return r.v
}

// Expect returns the holding value. If `r` is holding an error, the function panics
// with a message of the format: `str: error`.
func (r Result[T]) Expect(str string) T {
	if r.e != nil {
		panic(fmt.Sprintf("%s: %s", str, r.e))
	}

	return r.v
}

// V returns the expected value.
func (r Result[T]) V() T {
	return r.v
}

// E returns the unexpected value.
func (r Result[T]) E() error {
	return r.e
}

// VE returns V and E.
func (r Result[T]) VE() (T, error) {
	return r.V(), r.E()
}

// Then is executed if Result is not holding an unexpected error.
func (r Result[T]) Then(fn func(T)) Result[T] {
	if r.e == nil {
		fn(r.v)
	}

	return r
}

// ThenE is executed if Result is not holding an unexpected error.
//
// The function must return en error. The error can be nil.
// The returned error will be set to the Result's error.
func (r Result[T]) ThenE(fn func(T) error) Result[T] {
	if r.e == nil {
		r.e = fn(r.v)
	}

	return r
}

// Else is executed if Result is holding an error.
func (r Result[T]) Else(fn func(error)) Result[T] {
	if r.e != nil {
		fn(r.e)
	}

	return r
}

// ElseE is executed if Result is holding an error.
//
// The function must return an error. The error can be nil.
// The returned error will be set to the Result's error.
func (r Result[T]) ElseE(fn func(error) error) Result[T] {
	if r.e != nil {
		r.e = fn(r.e)
	}

	return r
}