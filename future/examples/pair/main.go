package main

import (
	"fmt"

	"github.com/dgrr/gtl/future"
)

var contents = []string{
	"hello",
	"testing",
	"world",
	"generics",
}

var prices = map[string]float64{
	"BTC": 57283.30,
	"ETH": 2874.17,
	"BNB": 620.2,
	"ADA": 0,
}

func getPrices() (vec gtl.Vec[gtl.Pair[string, gtl.Optional[float64]]]) {
	for k, v := range prices {
		vec.Append(
			gtl.NewPair(k, gtl.OptFromBool(v, v != 0)))
	}

	return
}

func main() {
	prices := getPrices()

	for _, pricePair := range prices {
		// TODO: this print panics bc the compiler fails to assert something
		fmt.Printf("%s: %0.2f\n", pricePair.First(), pricePair.Second().Or(0).Unwrap())
	}
}
