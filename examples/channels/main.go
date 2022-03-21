package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/dgrr/gtl"
)

func main() {
	ch := make(chan int, 128)

	s := gtl.MakeSenderChan(ch)
	r := gtl.MakeReceiverChan(ch)

	go func() {
		sendRandom(s)
		close(ch)
	}()

	r.Range(func(n int) {
		fmt.Println("recv", n)
	})
}

func sendRandom(s gtl.Sender[int]) {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 16; i++ {
		s.Send(rand.Int())
	}
}
