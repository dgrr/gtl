package gtl

type Sender[T any] struct {
	ch   chan T
	done chan struct{}
}

func MakeSender[T any](size int) Sender[T] {
	return MakeSenderChan(make(chan T, size))
}

func MakeSenderChan[T any](ch chan T) Sender[T] {
	return Sender[T]{
		ch:   ch,
		done: make(chan struct{}, 1),
	}
}

func (s Sender[T]) Get() chan T {
	return s.ch
}

func (s Sender[T]) Send(data T) bool {
	select {
	case s.ch <- data:
		return true
	case <-s.done:
		return false
	}
}

func (s Sender[T]) Close() error {
	close(s.ch)
	return nil
}

type Receiver[T any] struct {
	ch chan T
}

func MakeReceiver[T any](size int) Receiver[T] {
	return MakeReceiverChan(make(chan T, size))
}

func MakeReceiverChan[T any](ch chan T) Receiver[T] {
	return Receiver[T]{
		ch: ch,
	}
}

func (r Receiver[T]) Next() Optional[T] {
	v, ok := <-r.ch
	return OptionalWithCond[T](v, ok)
}

func (r Receiver[T]) Range(fn func(T)) {
	for v := range r.ch {
		fn(v)
	}
}

func (r Receiver[T]) RangeBool(fn func(T) bool) {
	for v := range r.ch {
		if !fn(v) {
			break
		}
	}
}

func (r Receiver[T]) Close() error {
	close(r.ch)
	return nil
}
