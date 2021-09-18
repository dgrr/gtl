package main

import (
	"fmt"

	"github.com/dgrr/gtl"
)

func main() {
	fmt.Println("Max between 20,30 is:", gtl.Max(20, 30))
	fmt.Println("Min between 20,30 is:", gtl.Min(20, 30))
}
