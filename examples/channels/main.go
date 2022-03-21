package main

import (
	"fmt"
	"io"
	"net"

	"github.com/dgrr/gtl"
)

func main() {
	ln, _ := net.Listen("tcp4", ":4444")

	for {
		c, err := ln.Accept()
		if err != nil {
			panic(err)
		}

		go handleConn(NewConn(c))
	}
}

func handleConn(c *Conn) {
	defer c.Close()

	b := make([]byte, 1024)

	for {
		n, err := c.c.Read(b)
		if err != nil {
			break
		}

		c.sender.Send(string(b[:n]))
	}

	fmt.Println("connection closed")
}

type Conn struct {
	c net.Conn

	ch gtl.Channel[string]

	sender gtl.Sender[string]
}

func (c *Conn) Close() error {
	c.ch.Close()
	return c.c.Close()
}

func NewConn(c net.Conn) *Conn {
	ch := gtl.MakeChan[string](128)

	sender, recv := ch.Split()

	cc := &Conn{
		c:      c,
		ch:     ch,
		sender: sender,
	}

	go cc.writeLoop(recv)

	return cc
}

func (c *Conn) writeLoop(recv gtl.Receiver[string]) {
	defer fmt.Println("write loop exit")

	recv.Range(func(data string) {
		c.c.Write([]byte(data))
	})
}

func (c *Conn) Write(data string) (int, error) {
	n := len(data)

	ok := c.sender.Send(data)
	if !ok {
		return -1, io.EOF
	}

	return n, nil
}
