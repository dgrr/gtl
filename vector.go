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

// NewVecSize creates a vector of type T with the given size and capacity.
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

// Del removes the element in the position of the iterator `it`.
//
// Returns true if the element has been removed.
func (vc *Vec[T]) Del(it Iterator[T]) (val T, erased bool) {
	nit, ok := it.(*Iter[T, int])
	if ok {
		return vc.DelByIndex(nit.Index())
	}

	return
}

// Filter filters the contents of Vec using cmpFn.
//
// If cmpFn returns true, the iterator will be removed.
func (vc *Vec[T]) Filter(cmpFn func(it Iterator[T]) bool) {
	for it := vc.iter(); it.Next(); {
		if !cmpFn(it) {
			vc.DelByIndex(it.Index())
			it.index = it.prev
		}
	}
}

// DelByIndex removes the element in the index `i`.
//
// Returns true if the element has been removed.
func (vc *Vec[T]) DelByIndex(i int) (val T, erased bool) {
	if vc.Len() <= i {
		return val, false
	}

	val = (*vc)[i]
	*vc = append((*vc)[:i], (*vc)[i+1:]...)

	return val, true
}

// Iter returns an iterator over the vector.
func (vc *Vec[T]) Iter() Iterator[T] {
	return vc.iter()
}

func (vc *Vec[T]) iter() *Iter[T, int] {
	it := &Iter[T, int]{
		v: nil,
		next: func(cnt int) (*T, int) {
			if cnt < vc.Len() {
				return &(*vc)[cnt], cnt + 1
			}

			return nil, cnt
		},
		advance: func(cnt, n int) (*T, int) {
			cnt += n
			if cnt < vc.Len() {
				return &(*vc)[cnt], cnt
			}

			return nil, cnt
		},
	}

	return it
}

// Index returns the index of an element inside the vector.
// The `cmpFn` lambda is used to perform the comparison. The lambda
// gets as input an Iterator[T] and should return true if the value matches the expected.
func (vc Vec[T]) Index(cmpFn func(it Iterator[T]) bool) int {
	i := -1
	for it := vc.iter(); it.Next(); {
		if cmpFn(it) {
			i = it.Index()
			break
		}
	}

	return i
}

func (vc Vec[T]) Contains(cmpFn func(it Iterator[T]) bool) bool {
	return vc.Index(cmpFn) >= 0
}

// Search iterates over the vector calling `cmpFn`.
//
// If cmpFn returns true, Search returns the current Iterator.
func (vc Vec[T]) Search(cmpFn func(v T) bool) Iterator[T] {
	for it := vc.iter(); it.Next(); {
		if cmpFn(it.Get()) {
			return it
		}
	}

	return nil
}

// SearchByValue is like Search but the input of `cmpFn` is the value, not the iterator.
func (vc Vec[T]) SearchByValue(cmpFn func(T) bool) Iterator[T] {
	for it := vc.iter(); it.Next(); {
		if cmpFn(it.Get()) {
			return it
		}
	}

	return nil
}
