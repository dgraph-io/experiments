package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/rpc"
)

type Query struct {
	d []byte
}

type Reply struct {
	d []byte
}

func setError(prev *error, n error) {
	if prev == nil {
		prev = &n
	}
}

type Worker struct {
}

func serveIt(conn io.ReadWriteCloser) {
	for {
		srv := &scodec{
			rwc:  conn,
			ebuf: bufio.NewWriter(conn),
		}
		rpc.ServeRequest(srv)
	}
}

func (w *Worker) Receive(query *Query, reply *Reply) error {
	fmt.Printf("Worker received: [%s]\n", string(query.d))
	reply.d = []byte("abcdefghij-Hello World!")
	return nil
}

func runServer(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Printf("listen(%q): %s\n", address, err)
		return err
	}
	fmt.Printf("Worker listening on %s\n", ln.Addr())
	go func() {
		for {
			cxn, err := ln.Accept()
			if err != nil {
				log.Fatalf("listen(%q): %s\n", address, err)
				return
			}
			log.Printf("Worker accepted connection to %s from %s\n",
				cxn.LocalAddr(), cxn.RemoteAddr())
			go serveIt(cxn)
		}
	}()
	return nil
}

func main() {
	addresses := map[int]string{
		1: "127.0.0.1:10000",
		2: "127.0.0.1:10001",
		3: "127.0.0.1:10002",
	}

	w := new(Worker)
	if err := rpc.Register(w); err != nil {
		log.Fatal(err)
	}

	for _, address := range addresses {
		if err := runServer(address); err != nil {
			log.Fatal(err)
		}
	}

	clients := make(map[int]*rpc.Client)
	for id, address := range addresses {
		conn, err := net.Dial("tcp", address)
		if err != nil {
			log.Fatal("dial", err)
		}
		cc := &ccodec{
			rwc: conn,
		}
		clients[id] = rpc.NewClientWithCodec(cc)
	}

	for i := 0; i < 10; i++ {
		/*
			client := clients[1]
			if client == nil {
				log.Fatal("Worker is nil")
			}
		*/

		for id, client := range clients {
			query := new(Query)
			query.d = []byte(fmt.Sprintf("id:%d Rand: %d", id, rand.Int()))
			reply := new(Reply)
			if err := client.Call("Worker.Receive", query, reply); err != nil {
				log.Fatal("call", err)
			}

			fmt.Printf("Returned: %s\n", string(reply.d))
		}
	}
}
