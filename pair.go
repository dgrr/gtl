package gtl

// Pair defines a pair of values.
type Pair[T, U any] struct {
	t T
	u U
}

// NewPair returns a Pair[T, U].
func NewPair[T, U any](t T, u U) Pair[T, U] {
	return Pair[T, U]{
		t: t,
		u: u,
	}
}

// Swap returns a Pair swapping both values.
func (p Pair[T, U]) Swap() Pair[U, T] {
	return NewPair(p.u, p.t)
}

// T returns the first value.
func (p Pair[T, U]) T() T {
	return p.t
}

func (p Pair[T, U]) First() T {
	return p.t
}

// U returns the second value.
func (p Pair[T, U]) U() U {
	return p.u
}

func (p Pair[T, U]) Second() U {
	return p.u
}

func (p Pair[T, U]) Both() (T, U) {
	return p.t, p.u
}

func (p Pair[T, U]) TU() (T, U) {
	return p.t, p.u
}
