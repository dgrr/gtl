package main

import (
	"fmt"
	"math/rand"
	"sort"

	. "github.com/dgrr/gtl"
)

func main() {
	var vec Vec[int]

	for i := 0; i < 20; i++ {
		vec.Append(128)
		vec.Append(
			rand.Intn(100))
	}

	vec.DelFn(func(v int) bool {
		return v == 128
	})

	sort.Slice(vec, func(i, j int) bool {
		return vec[i] < vec[j]
	})

	for vec.Len() != 0 {
		fmt.Println(
			vec.PopFront())
	}

	vec.Push(-20)

	fmt.Println("last value", vec.PopFront())
}
