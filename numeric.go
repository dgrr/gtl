package gtl

import "golang.org/x/exp/constraints"

// Max returns the max of the 2 passed values.
func Max[T constraints.Ordered](a, b T) (r T) {
	if a < b {
		r = b
	} else {
		r = a
	}

	return
}

// Min returns the min of the 2 passed values.
func Min[T constraints.Ordered](a, b T) (r T) {
	if a < b {
		r = a
	} else {
		r = b
	}

	return
}
