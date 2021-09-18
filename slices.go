package gtl

// Contains returns whether `vs` contains the element `e`.
func Contains[T comparable](vs []T, e T) bool {
	for _, v := range vs {
		if v == e {
			return true
		}
	}

	return false
}

// ExtractFrom a nested object of type E from type T.
func ExtractFrom[T, E any](set []T, fn func(T) E) []E {
	r := make([]E, len(set))
	for i := range set {
		r[i] = fn(set[i])
	}

	return r
}

// Filter iterates over `set` and gets the values that match `criteria`.
func Filter[T any](set []T, criteria func(T) bool) []T {
	r := make([]T, 0)
	for i := range set {
		if criteria(set[i]) {
			r = append(r, set[i])
		}
	}

	return r
}
