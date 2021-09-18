package gtl

// Optional defines an optional object.
type Optional[T any] struct {
	v *T
}

// NewOptional returns a new Optional value.
//
// If `v` is nil, Optional will be invalid.
func NewOptional[T any](v *T) Optional[T] {
	return Optional[T]{
		v: v,
	}
}

// OptionalFrom ...
func OptionalFrom[T any](v T, err error) (opt Optional[T]) {
	return opt.From(v, err)
}

// From ...
func (opt Optional[T]) From(v T, err error) Optional[T] {
	if err == nil {
		opt.v = &v
	}

	return opt
}

// Unwrap unwraps the value held.
func (opt Optional[T]) Unwrap() T {
	return opt.V()
}

// V returns the held value if it has been defined previously.
func (opt Optional[T]) V() (v T) {
	if opt.v != nil {
		v = *opt.v
	}

	return
}

// Or assigns `v` to `opt` only if the Optional value is not defined.
func (opt Optional[T]) Or(v T) Optional[T] {
	if opt.v == nil {
		opt.Set(v)
	}

	return opt
}

// IsOk returns true if Optional is holding a value.
func (opt Optional[T]) IsOk() bool {
	return opt.v != nil
}

// Set sets the value to the optional struct.
func (opt *Optional[T]) Set(v T) {
	if opt.v == nil {
		opt.v = new(T)
	}

	*opt.v = v
}

// Drop drops any previously set value.
func (opt *Optional[T]) Drop() {
	opt.v = nil
}
