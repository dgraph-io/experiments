package vrpc

import (
	"net"
	"net/rpc"
	"testing"

	"github.com/valyala/gorpc"
)

func BenchmarkPingPong_valyala(b *testing.B) {
	gorpc.RegisterType(&PostingList{})

	s := gorpc.NewTCPServer(":12345", PingPong)
	if err := s.Start(); err != nil {
		b.Fatal("While starting server on port 12345")
		return
	}
	defer s.Stop()

	req := NewRequest()
	c := gorpc.NewTCPClient(":12345")
	c.Start()
	defer c.Stop()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := c.Call(req)
		if err != nil {
			b.Fatalf("While running request: %v", err)
			return
		}
	}
}

func getTcpPipe(b *testing.B) (net.Conn, net.Conn) {
	ln, err := net.Listen("tcp", "127.0.0.1:12345")
	if err != nil {
		b.Fatalf("cannot listen to socket: %s", err)
	}

	ch := make(chan net.Conn, 1)
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			b.Fatalf("cannot accept incoming tcp conn: %s", err)
		}
		ch <- conn
	}()

	addr := ln.Addr().String()
	connC, err := net.Dial("tcp", addr)
	if err != nil {
		b.Fatalf("cannot dial %s: %s", addr, err)
	}
	connS := <-ch
	return connC, connS
}

func BenchmarkPingPong_netrpc(b *testing.B) {
	connC, connS := getTcpPipe(b)
	defer connC.Close()
	defer connS.Close()

	s := rpc.NewServer()
	if err := s.Register(&PostingList{}); err != nil {
		b.Fatalf("Error when registering rpc server: %v", err)
		return
	}
	go s.ServeConn(connS)

	c := rpc.NewClient(connC)
	defer c.Close()

	req := NewRequest()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var reply PostingList
		if err := c.Call("PostingList.PingPong", req, &reply); err != nil {
			b.Fatalf("While running request: %v", err)
			return
		}
	}
}
