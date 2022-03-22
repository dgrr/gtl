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
func OptionalFrom[T any](v T) (opt Optional[T]) {
	return opt.From(v)
}

// OptionalWithCond creates an optional using the second argument as condition.
// If the condition is true, the first argument is set as the underlying optional value.
func OptionalWithCond[T any](v T, cond bool) (opt Optional[T]) {
	return opt.WithCond(v, cond)
}

func (opt Optional[T]) WithCond(v T, cond bool) Optional[T] {
	if cond {
		opt.Set(v)
	}

	return opt
}

// From ...
func (opt Optional[T]) From(v T) Optional[T] {
	opt.Set(v)
	return opt
}

// Ptr returns a pointer to the value. If there's no value set, nil is returned.
func (opt Optional[T]) Ptr() *T {
	if opt.hasValue {
		return &opt.value
	}

	return nil
}

// Unwrap unwraps the value held.
func (opt Optional[T]) Unwrap() T {
	return opt.Get()
}

// Get returns the value held by the Optional.
func (opt Optional[T]) Get() T {
	return opt.Value()
}

// Value returns the value held by the Optional.
func (opt Optional[T]) Value() T {
	return opt.value
}

// Or assigns `v` to `opt` only if the Optional value is not defined.
func (opt Optional[T]) Or(v T) Optional[T] {
	if !opt.hasValue {
		opt.Set(v)
	}

	return opt
}

// Then will be called if Optional is holding a value.
func (opt Optional[T]) Then(fn func(T)) Optional[T] {
	if opt.hasValue {
		fn(opt.value)
	}
	return opt
}

// Else will be called if Optional is not holding any value.
func (opt Optional[T]) Else(fn func()) Optional[T] {
	if !opt.hasValue {
		fn()
	}
	return opt
}

// OrFn operates like Or but with the difference that OrFn will assign the object
// returned from a function.
//
// This function is useful if the alternative object is expensive to make. Instead of
// creating the object regardless the previous result, OrFn will only build the object
// if there's no previous value available.
func (opt Optional[T]) OrFn(fn func() T) Optional[T] {
	if !opt.hasValue {
		opt.Set(fn())
	}

	return opt
}

// IsOk returns true if Optional is holding a value.
func (opt Optional[T]) IsOk() bool {
	return opt.hasValue
}

// HasValue returns true if the Optional is holding a valid value.
func (opt Optional[T]) HasValue() bool {
	return opt.IsOk()
}

// Set sets the value to the optional struct.
func (opt *Optional[T]) Set(v T) {
	opt.value = v
	opt.hasValue = true
}

// Reset resets the value of the optional.
func (opt *Optional[T]) Reset() {
	opt.hasValue = false
}
