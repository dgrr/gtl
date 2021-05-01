package gtl

import "sync"

// Locker wraps `T` with a sync.Locker.
type Locker[T any] struct {
	sync.Mutex
	v T
}

// NewLocker returns a NewLocker.
func NewLocker[T any](vs ...T) *Locker[T] {
	var v T
	if len(vs) != 0 {
		v = vs[0]
	}

	return &Locker[T]{
		v: v,
	}
}

// V returns the value of T.
func (lck *Locker[T]) V() T {
	return lck.v
}

// Ptr returns a pointer to `T`.
func (lck *Locker[T]) Ptr() *T {
	return &lck.v
}

// Set sets the value wrapping the assign with the Lock and Unlock calls.
func (lck *Locker[T]) Set(v T) {
	lck.Lock()
	lck.v = v
	lck.Unlock()
}
