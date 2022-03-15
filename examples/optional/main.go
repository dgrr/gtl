package main

import (
	"fmt"
	"strconv"

	. "github.com/dgrr/gtl"
)

func myHandler(n Optional[int64]) {
	if n.IsOk() {
		fmt.Println("Handler recv", n.Get())
	}
}

func main() {
	n := OptionalFrom(
		strconv.ParseInt("123", 10, 64))
	myHandler(n)

	n.Drop()

	n = OptionalFrom(
		strconv.ParseInt("abc", 10, 64))
	myHandler(n)

	n = OptionalFrom(
		strconv.ParseInt("981a", 10, 64)).Or(0)
	myHandler(n)
}
