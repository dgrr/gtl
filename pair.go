package gtl

// Pair defines a pair of values.
type Pair[T, U any] struct {
	t T
	u U
}

// MakePair returns a Pair[T, U].
func MakePair[T, U any](t T, u U) Pair[T, U] {
	return Pair[T, U]{
		t: t,
		u: u,
	}
}

// Swap returns a Pair swapping both values.
func (p Pair[T, U]) Swap() Pair[U, T] {
	return MakePair(p.u, p.t)
}

// First returns the first element from the Pair.
func (p Pair[T, U]) First() T {
	return p.t
}

// Second returns the second element from the pair.
func (p Pair[T, U]) Second() U {
	return p.u
}

// Both returns both elements from the pair.
func (p Pair[T, U]) Both() (T, U) {
	return p.t, p.u
}
