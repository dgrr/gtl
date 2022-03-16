package gtl

import (
	"golang.org/x/exp/constraints"
	"time"
)

type orderedCriteria[K constraints.Ordered, T any] struct {
	key    K
	values []T
}

type Series[T, F any, X constraints.Ordered] struct {
	data []T
	fild []orderedCriteria[X, T]

	opts SeriesOpts[T, T, X]
}

type SeriesOpts[T, RX any, AX constraints.Ordered] struct {
	winTime func(T) bool

	filters []FilterFunc[T]
	aggs    []AggregateFunc[T, AX]
	rfs     []ResultFunc[T, RX]
}

func NewSeries[T any, X constraints.Ordered](opts SeriesOpts[T, T, X]) *Series[T, T, X] {
	return &Series[T, T, X]{
		opts: opts,
	}
}

func (s *Series[T, F, X]) Push(v T) {
	s.tryDelete()

	s.data = append(s.data, v)

	s.applyAll()
}

func (s *Series[T, F, X]) tryDelete() {
	i := 0

	for len(s.data) > i && s.opts.winTime(s.data[i]) {
		i++
	}

	if i != 0 {
		s.data = append(s.data[:0], s.data[i:]...)
	}
}

func (s *Series[T, F, X]) applyAll() {
	s.fild = nil

	wanted := make([]bool, len(s.data))

	for _, filter := range s.opts.filters {
		for i := range s.data {
			// if we want to remove it, then set it to true
			wanted[i] = wanted[i] || !filter(s.data[i])
		}
	}

	for _, agg := range s.opts.aggs {
		for i := range s.data {
			if wanted[i] {
				continue
			}

			var e *orderedCriteria[X, T]

			key := agg(s.data[i])

			for i := range s.fild {
				if s.fild[i].key == key {
					e = &s.fild[i]
				}
			}
			if e == nil {
				e = &orderedCriteria[X, T]{
					key: key,
				}
			}

			e.values = append(e.values, s.data[i])
		}
	}
}

type (
	WindowTimeFn[T any]                         func(T) time.Time
	FilterFunc[T any]                           func(T) bool
	AggregateFunc[T any, X constraints.Ordered] func(T) X
	ResultFunc[T, X any]                        func(X, T) X
)

func NewSeriesOpts[T, RX any, AX constraints.Ordered]() SeriesOpts[T, RX, AX] {
	return SeriesOpts[T, RX, AX]{}
}

func (s SeriesOpts[T, RX, AX]) WindowTime(d time.Duration, fn WindowTimeFn[T]) SeriesOpts[T, RX, AX] {
	s.winTime = func(x T) bool {
		return time.Since(fn(x)) > d
	}
	return s
}

func (s SeriesOpts[T, RX, AX]) Filter(fn FilterFunc[T]) SeriesOpts[T, RX, AX] {
	s.filters = append(s.filters, fn)
	return s
}

func (s SeriesOpts[T, RX, AX]) Aggregate(agg AggregateFunc[T, AX]) SeriesOpts[T, RX, AX] {
	s.aggs = append(s.aggs, agg)
	return s
}

func (s SeriesOpts[T, RX, AX]) Result(r ResultFunc[T, RX]) SeriesOpts[T, RX, AX] {
	s.rfs = append(s.rfs, r)
	return s
}
