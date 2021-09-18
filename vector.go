package gtl

// Vec represents a slice of type `T`.
type Vec[T any] []T

// NewVec returns a new Vec.
//
// If `elmnts` is defined, the contents will be append.
func NewVec[T any](elmnts ...T) Vec[T] {
	vc := Vec[T]{}
	vc.Append(elmnts...)

	return vc
}

func NewVecSize[T any](size, capacity int) Vec[T] {
	return (Vec[T])(make([]T, size, capacity))
}

// Get returns the element in the position `i`.
func (vc *Vec[T]) Get(i int) T {
	return (*vc)[i]
}

// Resize increases (if needed) the length of the vector to match `n`.
func (vc *Vec[T]) Resize(n int) {
	vc.Reserve(n)
	*vc = (*vc)[:n]
}

// Reserve ensures `vc` has at least len = n.
// It doesn't change the length if the len is less than `n`.
func (vc *Vec[T]) Reserve(n int) {
	if nSize := n - cap(*vc); nSize > 0 {
		*vc = append((*vc)[:cap(*vc)], make([]T, nSize)...)
	}
}

// Append appends `elmnts` to the end of the vector.
func (vc *Vec[T]) Append(elmnts ...T) {
	*vc = append(*vc, elmnts...)
}

// Push pushes the `elmnts` into the first positions of the vector.
func (vc *Vec[T]) Push(elmnts ...T) {
	*vc = append((*vc)[:len(elmnts)], *vc...)
	copy(*vc, elmnts)
}

// Front returns the first element of the vector.
func (vc *Vec[T]) Front() T {
	return (*vc)[0]
}

// FrontPtr returns a pointer to the first element of the vector.
func (vc *Vec[T]) FrontPtr() *T {
	return &(*vc)[0]
}

// Back returns the last element of the vector.
func (vc *Vec[T]) Back() T {
	return (*vc)[vc.Len()-1]
}

// BackPtr returns a pointer to the last element of the vector.
func (vc *Vec[T]) BackPtr() *T {
	return &(*vc)[vc.Len()-1]
}

// PopBack returns the last element and removes it from the vector.
func (vc *Vec[T]) PopBack() (e T) {
	if len(*vc) > 0 {
		e = (*vc)[len(*vc)-1]
		*vc = (*vc)[:len(*vc)-1]
	}

	return
}

// PopFront returns the first elements and removes it from the vector.
func (vc *Vec[T]) PopFront() (e T) {
	if len(*vc) > 0 {
		e = (*vc)[0]
		*vc = append((*vc)[:0], (*vc)[1:]...)
	}

	return
}

// Len returns the number of elements in `vc`.
func (vc *Vec[T]) Len() int {
	return len(*vc)
}

// Cap returns the capacity of the vector.
func (vc *Vec[T]) Cap() int {
	return cap(*vc)
}

// DelIndex removes the element in the index `i`.
//
// Returns true if the element has been removed.
func (vc *Vec[T]) DelIndex(i int) bool {
	if vc.Len() <= i {
		return false
	}

	*vc = append((*vc)[:i], (*vc)[i+1:]...)

	return true
}

func (vc *Vec[T]) DelFn(fn func(T) bool) (n int) {
	for i := 0; i < len(*vc); i++ {
		if fn((*vc)[i]) && vc.DelIndex(i) {
			n++
			i--
		}
	}

	return n
}

// Slice returns a slice from [start, end) of type T.
func (vc *Vec[T]) Slice(start, end int) []T {
	if end < start {
		end = start
	}

	return (*vc)[start:end]
}
