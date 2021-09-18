package main

import (
	"bytes"
	"fmt"

	"github.com/dgrr/gtl/future"
)

func main() {
	data := bytes.NewBuffer(
		[]byte("Hello world\r\n1234"))

	var b gtl.Bytes
	fmt.Println("len", b.Len())

	// TODO: complains here because the compiler says there are 2 functions called Resize, in Bytes and Vec.
	// If you remove the Bytes function, then it complains that no function called Resize exists.
	b.Resize(32)

	fmt.Println("len after resize", b.Len())

	_, err := b.ReadFrom(data)
	if err != nil {
		panic(err)
	}

	fmt.Println(b.Index('0'))
	fmt.Println(
		b[:b.Index('\r')].String())

	b = gtl.BytesFrom(data.Bytes())
	fmt.Printf("%s\n", b)
}
