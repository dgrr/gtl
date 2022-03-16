package gtl

type element[T any] struct {
	data T
	next *element[T]
}

type Queue[T any] struct {
	first *element[T]
	last  *element[T]
}

func (q *Queue[T]) Reset() {
	q.first = nil
	q.last = nil
}

func (q *Queue[T]) Front() (v Optional[T]) {
	if q.first != nil {
		v.Set(q.first.data)
	}

	return
}

func (q *Queue[T]) PushFront(data T) {
	if q.first == nil {
		q.first = &element[T]{
			data: data,
		}
		q.last = q.first
	} else {
		e := &element[T]{
			data: data,
			next: q.first,
		}
		q.first = e
	}
}

func (q *Queue[T]) PushBack(data T) {
	if q.first == nil {
		q.first = &element[T]{
			data: data,
		}
		q.last = q.first
	} else {
		q.last.next = &element[T]{
			data: data,
		}
		q.last = q.last.next
	}
}

func (q *Queue[T]) Pop() (v Optional[T]) {
	if q.first != nil {
		v.Set(q.first.data)
		q.first = q.first.next
	}

	return
}
