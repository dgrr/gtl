package gtl

type Ordered interface {
	int | int8 | int16 | int32 | int64 |
		uint | uint8 | uint16 | uint32 | uint64 | uintptr |
		float32 | float64 | string
}

// Max returns the max of the 2 passed values.
func Max[T Ordered](a, b T) (r T) {
	if a < b {
		r = b
	} else {
		r = a
	}

	return
}

// Min returns the min of the 2 passed values.
func Min[T Ordered](a, b T) (r T) {
	if a < b {
		r = a
	} else {
		r = b
	}

	return
}
