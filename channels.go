package gtl

import (
	"sync/atomic"
)

type Sender[T any] struct {
	ch chan<- T
}

func MakeSender[T any](ch chan<- T) Sender[T] {
	return Sender[T]{
		ch: ch,
	}
}

func (s Sender[T]) Get() chan<- T {
	return s.ch
}

func (s Sender[T]) Send(data T) bool {
	select {
	case s.ch <- data:
		return true
	default:
		return false
	}
}

func (s Sender[T]) Len() int {
	return len(s.ch)
}

func (s Sender[T]) Close() error {
	close(s.ch)
	return nil
}

type Receiver[T any] struct {
	ch <-chan T
}

func MakeReceiver[T any](ch <-chan T) Receiver[T] {
	return Receiver[T]{
		ch: ch,
	}
}

func (r Receiver[T]) Next() Optional[T] {
	v, ok := <-r.ch
	return OptionalWithCond[T](v, ok)
}

func (r Receiver[T]) Get() <-chan T {
	return r.ch
}

func (r Receiver[T]) Range(fn func(T)) {
	for v := range r.ch {
		fn(v)
	}
}

func (r Receiver[T]) Len() int {
	return len(r.ch)
}

func (r Receiver[T]) RangeBool(fn func(T) bool) {
	for v := range r.ch {
		if !fn(v) {
			break
		}
	}
}

type Channel[T any] struct {
	ch     chan T
	closed int32
}

func MakeChan[T any](size int) Channel[T] {
	return Channel[T]{
		ch: make(chan T, size),
	}
}

func (ch Channel[T]) Len() int {
	return len(ch.ch)
}

func (ch Channel[T]) Split() (Sender[T], Receiver[T]) {
	return MakeSender(ch.ch), MakeReceiver(ch.ch)
}

func (ch *Channel[T]) Close() error {
	if atomic.CompareAndSwapInt32(&ch.closed, 0, 1) {
		close(ch.ch)
	}

	return nil
}
