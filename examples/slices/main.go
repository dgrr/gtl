package main

import (
	"fmt"

	"github.com/dgrr/gtl"
)

func main() {
	v := []int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Printf("%v contains %d ? %v\n", v, 20, gtl.Contains(v, 20))
	fmt.Printf("%v contains %d ? %v\n", v, 4, gtl.Contains(v, 4))

	v = gtl.FilterInPlace(v, func(e int) bool {
		return e < 4
	})

	fmt.Printf("%v\n", v)
}
