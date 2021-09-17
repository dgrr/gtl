package gtl

type element[T any] struct {
	v T
	next *element[T]
}

// List defines a linked-list.
type List[T any] struct {
	next *element[T]
}

// Add adds v to the front of the linked list.
func (lst *List[T]) Add(v T) {
	e := element[T]{
		v: v,
		next: lst.next,
	}
	lst.next = &e
}

// PopFront returns the first element removing it from the linked list.
func (lst *List[T]) PopFront() (v T) {
	if lst.next != nil {
		v = lst.next.v
		lst.next = lst.next.next
	}

	return
}

// Iter returns an iterator for the linked list.
func (lst *List[T]) Iter() Iterator[T] {
	iter := &Iter[T, *element[T]]{
		index: lst.next,
		next: func(prev *element[T]) (*T, *element[T]) {
			if prev == nil {
				return nil, nil
			}

			return &prev.v, prev.next
		},
	}

	return iter
}