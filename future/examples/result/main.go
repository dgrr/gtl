package main

import (
	"errors"
	"fmt"
	"strconv"

	. "github.com/dgrr/gtl/future"
)

func AddIfEven(a, b int) (r Result[int]) {
	if a%2 != 0 || b%2 != 0 {
		return r.Err(
			errors.New("a or b is an odd number"))
	}

	return r.Ok(a + b)
}

func main() {
	s := "19oo"

	r := NewResult(
		strconv.Atoi(s)).Or(-1)
	if r.V() == -1 {
		fmt.Printf("%s conversion failed\n", s)
	}

	if r := AddIfEven(2, 1); r.E() != nil {
		fmt.Printf("error: %s\n", r.E())
	}

	AddIfEven(4, 2).Then(func(res int) {
		fmt.Println("Res:", res)
	})

	fmt.Println(
		"adding odd numbers", AddIfEven(1, 2).Or(-1).V())
}
