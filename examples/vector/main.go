package main

import (
	"fmt"
	"math/rand"
	"sort"

	"github.com/dgrr/gtl"
)

func main() {
	var vec gtl.Vec[int]

	for i := 0; i < 20; i++ {
		vec.Append(128)
		vec.Append(
			rand.Intn(100))
	}

	vec.Filter(func(it gtl.Iterator[int]) bool {
		return it.Get() != 128
	})

	sort.Slice(vec, func(i, j int) bool {
		return vec[i] < vec[j]
	})

	it := vec.Iter()
	for i := 0; i < 6; i++ {
		it.Next()
	}

	if e, ok := vec.Del(it); ok {
		fmt.Println("erased", e)
	}

	for vec.Len() != 0 {
		fmt.Println(
			vec.PopFront())
	}

	vec.Push(-20)

	fmt.Println("last value", vec.PopFront())
}
