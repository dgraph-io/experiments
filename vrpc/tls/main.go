package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"math/rand"
	"net/rpc"

	"github.com/Sirupsen/logrus"
)

var log = logrus.WithField("pkg", "tls")

type PostingList struct {
	Uids []uint64
}

func NewRequest() *PostingList {
	req := new(PostingList)
	sz := 250000
	req.Uids = make([]uint64, sz)
	// Generate a 2MB request.
	for i := 0; i < sz; i++ {
		req.Uids[i] = uint64(rand.Int63())
	}
	return req
}

func (p *PostingList) PingPong(req *PostingList, reply *PostingList) error {
	reply.Uids = make([]uint64, len(req.Uids))
	for i := 0; i < len(req.Uids); i++ {
		reply.Uids[i] = req.Uids[i]
	}

	return nil
}

var addr = "127.0.0.1:12345"

func runServer(ch chan bool, done chan bool) {
	cert, err := tls.LoadX509KeyPair("../cert.pem", "../key.pem")
	if err != nil {
		log.Fatalf("While loading tls certs: %v", err)
		return
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", addr, &config)
	if err != nil {
		log.Fatalf("When listening: %v", err)
		return
	}
	s := rpc.NewServer()
	if err := s.Register(&PostingList{}); err != nil {
		log.Fatalf("Error when registering rpc server: %v", err)
		return
	}

	ch <- true
	log.Debugln("Ready to accept new connection")
	conn, err := ln.Accept()
	if err != nil {
		log.Fatalf("cannot accept incoming tcp conn: %s", err)
		return
	}
	defer conn.Close()
	log.Debugln("Accepted a connection")
	go s.ServeConn(conn)
	<-done
}

func getClientConn(ch chan bool) *tls.Conn {
	<-ch // Block until server is ready.
	ca_pool := x509.NewCertPool()
	scert, err := ioutil.ReadFile("../cert.pem")
	if err != nil {
		log.Fatalf("While reading cert.pem: %v", err)
	}
	ca_pool.AppendCertsFromPEM(scert)
	cconf := tls.Config{RootCAs: ca_pool, InsecureSkipVerify: true}

	/*
		insConn, err := net.DialTimeout("tcp", addr, 10*time.Second)
		if err != nil {
			log.Fatalf("While dialing via net: %v", err)
		}
		log.Debugln("Got connection via net.Dial")
	*/

	connC, err := tls.Dial("tcp", addr, &cconf)
	// connC := tls.Client(insConn, &cconf)
	log.Debugln("Converted to tls")
	if err := connC.Handshake(); err != nil {
		log.Fatalf("While handshaking: %v", err)
	}
	log.Debugln("Shook hands")
	return connC
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	ch := make(chan bool)
	done := make(chan bool)
	go runServer(ch, done)
	connC := getClientConn(ch)
	defer connC.Close()

	c := rpc.NewClient(connC)
	defer c.Close()

	req := NewRequest()
	var reply PostingList
	if err := c.Call("PostingList.PingPong", req, &reply); err != nil {
		log.Fatalf("While running request: %v", err)
		return
	}
	log.Printf("Got reply of len: %v", len(reply.Uids))
	done <- true
}
