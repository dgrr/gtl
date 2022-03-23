package gtl

// Contains returns whether `vs` contains the element `e` by comparing vs[i] == e.
func Contains[T comparable](vs []T, e T) bool {
	for _, v := range vs {
		if v == e {
			return true
		}
	}

	return false
}

// ExtractFrom extracts a nested object of type E from type T.
//
// This function is useful if we have a set of type `T` nad we want to
// extract the type E from any T.
func ExtractFrom[T, E any](set []T, fn func(T) E) []E {
	r := make([]E, len(set))
	for i := range set {
		r[i] = fn(set[i])
	}

	return r
}

// Filter iterates over `set` and gets the values that match `criteria`.
//
// Filter will return a new allocated slice.
func Filter[T any](set []T, criteria func(T) bool) []T {
	r := make([]T, 0)
	for i := range set {
		if criteria(set[i]) {
			r = append(r, set[i])
		}
	}

	return r
}

// FilterInPlace filters the contents of `set` using `criteria`.
//
// FilterInPlace returns `set`.
func FilterInPlace[T any](set []T, criteria func(T) bool) []T {
	for i := 0; i < len(set); i++ {
		if !criteria(set[i]) {
			set = append(set[:i], set[i+1:]...)
			i--
		}
	}

	return set
}

// Delete the first occurrence of a type from a set.
func Delete[T comparable](set []T, value T) []T {
	for i := 0; i < len(set); i++ {
		if set[i] == value {
			set = append(set[:i], set[i:]...)
			break
		}
	}

	return set
}

// DeleteAll occurrences from a set.
func DeleteAll[T comparable](set []T, value T) []T {
	for i := 0; i < len(set); i++ {
		if set[i] == value {
			set = append(set[:i], set[i:]...)
			i--
		}
	}

	return set
}
