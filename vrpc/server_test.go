// vrpc package benchmarks the various RPC libraries.
// 1. A custom tcp library by valyala (custom TCP)
// 2. net/rpc standard library (TCP)
// 3. crypto/tls over net/rpc standard libraries. (TLS over TCP)
//
// The results as of today, on my desktop are these:
// BenchmarkPingPong_2MB_valyala	       2	 554987424 ns/op
// BenchmarkPingPong_2MB_tlsrpc	      20	  77509099 ns/op
// BenchmarkPingPong_2MB_netrpc	     100	  17450330 ns/op
// BenchmarkPingPong_2KB_valyala	   10000	    167166 ns/op
// BenchmarkPingPong_2KB_tlsrpc	   10000	    176208 ns/op
// BenchmarkPingPong_2KB_netrpc	   10000	    102846 ns/op
//
// So, valyala is consistently slow.
// TLS-TCP vs only TCP ranges from 1.7x to 4.5x performance
// penalty, which is significant.
// For now, it probably makes sense to stick to just TCP,
// and see if we need to worry about encrypting inter-node connections later.

package vrpc

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"
	"net/rpc"
	"testing"
	"time"

	"github.com/valyala/gorpc"
)

var mb2 = 250000
var kb2 = 250

// Benchmark valyala/gorpc TCP connection for 2MB payload.
func BenchmarkPingPong_2MB_valyala(b *testing.B) {
	gorpc.RegisterType(&PostingList{})

	s := gorpc.NewTCPServer(":12345", PingPong)
	if err := s.Start(); err != nil {
		b.Fatal("While starting server on port 12345")
		return
	}
	defer s.Stop()

	req := NewRequest(mb2)
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

func BenchmarkPingPong_2MB_tlsrpc(b *testing.B) {
	ready := make(chan bool)
	done := make(chan bool)
	go runTlsServer(b, ready, done)
	connC := getTlsClientConn(b, ready)
	defer connC.Close()

	c := rpc.NewClient(connC)
	defer c.Close()

	req := NewRequest(mb2)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var reply PostingList
		if err := c.Call("PostingList.PingPong", req, &reply); err != nil {
			b.Fatalf("While running request: %v", err)
			return
		}
	}
	b.StopTimer()
	done <- true
}

func BenchmarkPingPong_2MB_netrpc(b *testing.B) {
	ready := make(chan bool)
	done := make(chan bool)
	addr := "127.0.0.1:12346"
	go runServer(b, addr, ready, done)

	<-ready // Block until server is ready.

	connC, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		b.Fatalf("cannot dial. Error: %v", err)
		return
	}
	defer connC.Close()

	c := rpc.NewClient(connC)
	defer c.Close()

	req := NewRequest(mb2)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var reply PostingList
		if err := c.Call("PostingList.PingPong", req, &reply); err != nil {
			b.Fatalf("While running request: %v", err)
			return
		}
	}
	b.StopTimer()
	done <- true
}

/*
func BenchmarkPingPong_2KB_udp(b *testing.B) {
	ready := make(chan bool)
	done := make(chan bool)
	addr := "127.0.0.1:12333"
	go runUdpServer(b, addr, ready, done)

	saddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		b.Fatalf("While resolving: %v", err)
		return
	}

	fmt.Println("Waiting for ready")
	<-ready
	fmt.Println("Server is now ready")
	connC, err := net.DialUDP("udp", nil, saddr)
	if err != nil {
		b.Fatalf("While dialing: %v", err)
		return
	}
	fmt.Println("Dial done")

	c := rpc.NewClient(connC)
	defer c.Close()

	w := new(bytes.Buffer)
	for i := 0; i < 1000; i++ {
		binary.Write(w, binary.LittleEndian, rand.Int63())
	}
	req := w.Bytes()
	fmt.Printf("Got buffer of size: %v\n", len(req))
	b.ResetTimer()

	for i := 0; i < 10; i++ {
		fmt.Println("Sending call")
		n, err := connC.Write(req)
		if err != nil {
			b.Fatalf("While writing: %v", err)
			return
		}
		fmt.Printf("Wrote bytes: %v\n", n)
	}
	b.StopTimer()
	connC.Write([]byte("0"))
	fmt.Println("Sending done")
	done <- true
}
*/

func BenchmarkPingPong_2KB_valyala(b *testing.B) {
	gorpc.RegisterType(&PostingList{})

	s := gorpc.NewTCPServer(":12345", PingPong)
	if err := s.Start(); err != nil {
		b.Fatal("While starting server on port 12345")
		return
	}
	defer s.Stop()

	req := NewRequest(kb2)
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

func BenchmarkPingPong_2KB_tlsrpc(b *testing.B) {
	ready := make(chan bool)
	done := make(chan bool)
	go runTlsServer(b, ready, done)
	connC := getTlsClientConn(b, ready)
	defer connC.Close()

	c := rpc.NewClient(connC)
	defer c.Close()

	req := NewRequest(kb2)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var reply PostingList
		if err := c.Call("PostingList.PingPong", req, &reply); err != nil {
			b.Fatalf("While running request: %v", err)
			return
		}
	}
	b.StopTimer()
	done <- true
}

func BenchmarkPingPong_2KB_netrpc(b *testing.B) {
	ready := make(chan bool)
	done := make(chan bool)
	addr := "127.0.0.1:12348"
	go runServer(b, addr, ready, done)

	<-ready // Block until server is ready.

	connC, err := net.DialTimeout("tcp", addr, 10*time.Second)
	if err != nil {
		b.Fatalf("cannot dial. Error: %v", err)
		return
	}
	defer connC.Close()

	c := rpc.NewClient(connC)
	defer c.Close()

	req := NewRequest(kb2)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		var reply PostingList
		if err := c.Call("PostingList.PingPong", req, &reply); err != nil {
			b.Fatalf("While running request: %v", err)
			return
		}
	}
	b.StopTimer()
	done <- true
}

func runServer(b *testing.B, addr string, ready, done chan bool) {
	ln, err := net.Listen("tcp", addr)
	if err != nil {
		b.Fatalf("Cannot listen to socket: %s", err)
		return
	}
	defer ln.Close()

	s := rpc.NewServer()
	if err := s.Register(&PostingList{}); err != nil {
		b.Fatalf("Error when registering rpc server: %v", err)
		return
	}

	ready <- true
	conn, err := ln.Accept()
	if err != nil {
		b.Fatalf("Cannot accept incoming: %v", err)
		return
	}
	defer conn.Close()

	go s.ServeConn(conn)
	<-done
}

/*
func runUdpServer(b *testing.B, addr string, ready, done chan bool) {
	saddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		b.Fatalf("While resolving: %v", err)
		return
	}

	conn, err := net.ListenUDP("udp", saddr)
	if err != nil {
		b.Fatalf("Cannot listen to socket: %s", err)
		return
	}
	defer conn.Close()
	fmt.Println("Listen UDP OK.")

	ready <- true
	fmt.Println("serving connection via udp")
	buf := make([]byte, 2<<20)
	go func() {
		for {
			fmt.Println("Waiting for UDP packets")
			n, addr, err := conn.ReadFromUDP(buf)
			if err != nil {
				b.Fatalf("While reading from udp: %v", err)
				return
			}
			fmt.Printf("Got n bytes: %v Addr: %v\n", n, addr)
			if n == 1 {
				break
			}
		}
	}()
	// go s.ServeConn(conn)
	fmt.Println("Blocking on done")
	<-done
	fmt.Println("I AM DONE")
}
*/

// Benchmark TLS over TCP connection for 2MB payload.
var tlsAddr = "127.0.0.1:12347"

func runTlsServer(b *testing.B, ready, done chan bool) {
	cert, err := tls.LoadX509KeyPair("./cert.pem", "./key.pem")
	if err != nil {
		b.Fatalf("While loading tls certs: %v", err)
		return
	}
	config := tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", tlsAddr, &config)
	if err != nil {
		b.Fatalf("When listening: %v", err)
		return
	}
	defer ln.Close()

	s := rpc.NewServer()
	if err := s.Register(&PostingList{}); err != nil {
		b.Fatalf("Error when registering rpc server: %v", err)
		return
	}

	ready <- true

	conn, err := ln.Accept()
	if err != nil {
		b.Fatalf("cannot accept incoming tcp conn: %s", err)
		return
	}
	defer conn.Close()
	go s.ServeConn(conn)
	<-done
}

func getTlsClientConn(b *testing.B, ready chan bool) *tls.Conn {
	<-ready
	ca_pool := x509.NewCertPool()
	scert, err := ioutil.ReadFile("./cert.pem")
	if err != nil {
		b.Fatalf("While reading cert.pem: %v", err)
	}
	ca_pool.AppendCertsFromPEM(scert)
	cconf := tls.Config{RootCAs: ca_pool, InsecureSkipVerify: true}

	insConn, err := net.DialTimeout("tcp", tlsAddr, 10*time.Second)
	if err != nil {
		b.Fatalf("While dialing via net: %v", err)
	}

	connC := tls.Client(insConn, &cconf)
	if err := connC.Handshake(); err != nil {
		b.Fatalf("While handshaking: %v", err)
	}
	return connC
}
