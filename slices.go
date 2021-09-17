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
