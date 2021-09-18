package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"os/signal"

	"github.com/dgrr/gtl"
)

func main() {
	ln := gtl.NewResult(
		net.Listen("tcp", ":42421"),
	).Expect("error listening")

	serve(ln)
}

func serve(ln net.Listener) error {
	s := Server{
		ln: ln,
	}

	go s.handleSignals()

	for {
		c := gtl.NewResult(
			ln.Accept(),
		).Expect("error accepting client")

		// TODO: obvious race condition here
		s.conns.Append(c)

		go s.handleConn(c)
	}

	return nil
}

type Server struct {
	ln    net.Listener
	conns gtl.Vec[net.Conn]
}

func (s *Server) handleSignals() {
	sch := make(chan os.Signal, 1)

	signal.Notify(sch, os.Interrupt)

	select {
	case <-sch:
	}

	signal.Stop(sch)
	close(sch)

	for _, c := range s.conns {
		c.Close()
	}

	s.ln.Close()
}

func (s *Server) handleConn(c net.Conn) {
	b := gtl.NewVecSize[byte](128, 128)

	for {
		n, err := c.Read(b)
		if err != nil {
			break
		}

		i := bytes.IndexByte(b, '\n')

		fmt.Printf("Recv: %s\n", b.Slice(0, i))

		_, err = c.Write(b[:n])
		if err != nil {
			break
		}
	}
}
