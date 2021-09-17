package gtl

type Comparable interface {
	Ordered
}

// Contains returns whether `vs` contains the element `e`.
func Contains[T Comparable](vs []T, e T) bool {
	for _, v := range vs {
		if v == e {
			return true
		}
	}

	return false
}

func Extract[T, E any](set []T, fn func(T) E) []E {
	r := make([]E, len(set))
	for i := range set {
		r[i] = fn(set[i])
	}

	return r
}
