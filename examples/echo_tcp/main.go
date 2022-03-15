package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"

	"github.com/dgrr/gtl"
)

func main() {
	gtl.NewResult(
		net.Listen("tcp", ":42421"),
	).Expect("error listening")
}

func serve(ln net.Listener) error {
	s := Server{
		ln: ln,
	}

	go s.handleSignals()

	for {
		c := gtl.NewResult(
			ln.Accept(),
		).ElseE(func(err error) error {
			if strings.Contains(err.Error(), "use of closed") {
				err = nil
			}

			return err
		}).Expect("error accepting client")

		if c == nil {
			break
		}

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

	for it := s.conns.Iter(); it.Next(); {
		it.V().Close()
	}

	s.ln.Close()
}

func (s *Server) handleConn(c net.Conn) {
	b := gtl.NewBytes(128, 128)

	defer func() {
		it := s.conns.Search(func(nc net.Conn) bool {
			return nc == c
		})
		if it != nil {
			remoteAddr := it.V().RemoteAddr()
			if _, ok := s.conns.Del(it); ok {
				log.Printf("Removed: %s\n", remoteAddr)
			}
		}
	}()

	for {
		n, err := b.ReadFrom(c)
		if err != nil {
			break
		}

		i := b.Index('\n')

		fmt.Printf("Recv: %s\n", b.Slice(0, i))

		_, err = b.LimitWriteTo(c, int(n))
		if err != nil {
			break
		}
	}
}
