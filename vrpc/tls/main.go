package main

import (
	"crypto/tls"
	"math/rand"
	"net"
	"net/rpc"
	"time"

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

func PingPong(addr string, req interface{}) interface{} {
	pl := req.(*PostingList)
	reply := new(PostingList)
	reply.Uids = make([]uint64, len(pl.Uids))
	for i := 0; i < len(pl.Uids); i++ {
		reply.Uids[i] = pl.Uids[i]
	}
	return reply
}

func (pl *PostingList) PingPong(req *PostingList, reply *PostingList) error {
	reply = PingPong("", req).(*PostingList)
	return nil
}

func getTlsTcpPipe() (net.Conn, net.Conn) {
	cert, err := tls.LoadX509KeyPair("../cert.pem", "../key.pem")
	if err != nil {
		log.Fatalf("While loading tls certs: %v", err)
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", "127.0.0.1:12347", &config)
	if err != nil {
		log.Fatalf("When listening: %v", err)
	}

	ch := make(chan net.Conn, 1)
	go func() {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalf("cannot accept incoming tcp conn: %s", err)
		}
		log.Debug("Accepted a connection")
		ch <- conn
	}()

	log.Debugln("Set tls to listen")

	addr := ln.Addr().String()
	log.Debugf("Got addr: %v", addr)
	/*
		ca_pool := x509.NewCertPool()
		scert, err := ioutil.ReadFile("../cert.pem")
		if err != nil {
			log.Fatalf("While reading cert.pem: %v", err)
		}
		ca_pool.AppendCertsFromPEM(scert)
		cconf := tls.Config{RootCAs: ca_pool}
	*/

	dialer := net.Dialer{Timeout: time.Second}
	connC, err := tls.DialWithDialer(&dialer, "tcp", addr, nil)
	if err != nil {
		log.Fatalf("cannot dial: %s: %s", addr, err)
	}
	log.Debugln("tls dial done")
	connS := <-ch
	log.Debugln("Returning two connections")
	return connC, connS
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	connC, connS := getTlsTcpPipe()
	defer connC.Close()
	defer connS.Close()

	s := rpc.NewServer()
	if err := s.Register(&PostingList{}); err != nil {
		log.Fatalf("Error when registering rpc server: %v", err)
		return
	}
	go s.ServeConn(connS)

	c := rpc.NewClient(connC)
	defer c.Close()

	req := NewRequest()
	var reply PostingList
	if err := c.Call("PostingList.PingPong", req, &reply); err != nil {
		log.Fatalf("While running request: %v", err)
		return
	}
}
