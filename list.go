package gtl

type listElement[T any] struct {
	v    T
	next *listElement[T]
}

// List defines a linked-list.
type List[T any] struct {
	next *listElement[T]
}

// Add adds v to the front of the linked list.
func (lst *List[T]) Add(v T) {
	e := listElement[T]{
		v:    v,
		next: lst.next,
	}
	lst.next = &e
}

// PopFront returns the first listElement removing it from the linked list.
func (lst *List[T]) PopFront() (v T) {
	if lst.next != nil {
		v = lst.next.v
		lst.next = lst.next.next
	}

	return
}

// Iter returns an iterator for the linked list.
func (lst *List[T]) Iter() Iterator[T] {
	iter := &Iter[T, *listElement[T]]{
		index: lst.next,
		next: func(prev *listElement[T]) (*T, *listElement[T]) {
			if prev == nil {
				return nil, nil
			}

			return &prev.v, prev.next
		},
	}

	return iter
}
