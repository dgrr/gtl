package gtl

import (
	"testing"
	"time"
)

func TakeTime(trade Trade) time.Time {
	return trade.Time
}

type Trade struct {
	Time  time.Time
	Price float64
	Qty   float64
}

func TestSeries(t *testing.T) {
	s := NewSeries(
		NewSeriesOpts[Trade, Trade, float64]().
			WindowTime(
				time.Second, TakeTime,
			).
			Filter(func(trade Trade) bool {
				return trade.Qty > 5
			}).
			Aggregate(func(trade Trade) float64 {
				return trade.Price
			}),
	)

	trades := []Trade{
		{
			Time:  time.Now(),
			Price: 100,
			Qty:   1,
		},
		{
			Time:  time.Now(),
			Price: 100,
			Qty:   5,
		},
	}

	for _, trade := range trades {
		s.Push(trade)
	}
}
