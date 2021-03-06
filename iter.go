package gtl

// Iterator defines an interface for iterative objects.
type Iterator[T any] interface {
	// Next increments the iterator.
	Next() bool
	// Advance advances the cursor `n` steps. Returns false if `n` overflows.
	Advance(n int) bool
	// Get returns the value held in the iterator.
	Get() T
	// Ptr returns a pointer to T.
	Ptr() *T
}

// Iter implements the Iterator[T] interface.
//
// The iterator might not be directly used by the user, but as a mean
// to iterate over gtl defined data structures.
//
// T defines the object held by the Iterator.
// T2 is any object that keeps the previous state in order to
// help the iterator to move forward. For example, a vector might use as T2
// an integer, to be used as index for accessing the next element.
type Iter[T, T2 any] struct {
	v       *T
	prev    T2
	index   T2
	next    func(T2) (*T, T2)
	advance func(T2, int) (*T, T2)
}

// Index returns the index used in the iterator.
func (it *Iter[T, T2]) Index() T2 {
	return it.prev
}

// Next advances the iterator pointer.
func (it *Iter[T, T2]) Next() bool {
	it.prev = it.index
	it.v, it.index = it.next(it.index)

	return it.v != nil
}

func (it *Iter[T, T2]) Advance(n int) bool {
	it.v, it.index = it.advance(it.index, n)

	return it.v != nil
}

// Get returns the value held in the iterator.
func (it *Iter[T, T2]) Get() T {
	return *it.v
}

// Ptr returns a pointer to the value held in the iterator.
func (it *Iter[T, T2]) Ptr() *T {
	return it.v
}
