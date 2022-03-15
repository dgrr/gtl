package gtl

// Optional defines an optional object.
type Optional[T any] struct {
	value    T
	hasValue bool
}

// MakeOptional returns a new Optional value.
//
// If `v` is nil, Optional will be invalid.
func MakeOptional[T any](v *T) (opt Optional[T]) {
	if v != nil {
		opt.Set(*v)
	}

	return opt
}

// OptionalFrom ...
func OptionalFrom[T any](v T, err error) (opt Optional[T]) {
	return opt.From(v, err)
}

func OptionalFromBool[T any](v T, cond bool) (opt Optional[T]) {
	if cond {
		opt.Set(v)
	}

	return
}

// From ...
func (opt Optional[T]) From(v T, err error) Optional[T] {
	if err == nil {
		opt.Set(v)
	}

	return opt
}

// Unwrap unwraps the value held.
func (opt Optional[T]) Unwrap() T {
	return opt.Get()
}

// V returns the held value if it has been defined previously.
func (opt Optional[T]) Get() T {
	return opt.value
}

// Or assigns `v` to `opt` only if the Optional value is not defined.
func (opt Optional[T]) Or(v T) Optional[T] {
	if !opt.hasValue {
		opt.value = v
		opt.hasValue = true
	}

	return opt
}

// IsOk returns true if Optional is holding a value.
func (opt Optional[T]) IsOk() bool {
	return opt.hasValue
}

// Set sets the value to the optional struct.
func (opt *Optional[T]) Set(v T) {
	opt.value = v
	opt.hasValue = true
}

// Drop drops any previously set value.
func (opt *Optional[T]) Drop() {
	opt.hasValue = false
}
