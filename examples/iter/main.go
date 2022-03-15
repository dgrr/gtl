package main

import (
	"fmt"
	"math/rand"

	"github.com/dgrr/gtl"
)

func main() {
	var list gtl.List[int]

	for i := 0; i < 20; i++ {
		list.Add(
			rand.Intn(100))
	}

	printAll(list.Iter())

	for i := 0; i < 20; i++ {
		list.PopFront()
	}
}

func printAll(it gtl.Iterator[int]) {
	for it.Next() {
		fmt.Println(it.Get())
	}
}
