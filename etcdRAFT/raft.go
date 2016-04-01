package main

import (
	"bytes"
	"encoding/gob"
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net"
	"net/rpc"
	"strconv"
	"time"

	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
	"github.com/dgraph-io/dgraph/conn"
	"github.com/dgraph-io/dgraph/x"
	"golang.org/x/net/context"
)

var (
	hb    = 1
	pools = make(map[uint64]*conn.Pool)
	addrs = make(map[uint64]string)
	glog  = x.Log("RAFT")
)

type node struct {
	id     uint64
	addr   string
	ctx    context.Context
	pstore map[string]string
	store  *raft.MemoryStorage
	cfg    *raft.Config
	raft   raft.Node
	ticker <-chan time.Time
	done   <-chan struct{}
}

type Worker struct {
}

type raftRPC struct {
	Ctx     context.Context
	Message raftpb.Message
}

type helloRPC struct {
	Id   uint64
	Addr string
}

func (w *Worker) Hello(query *conn.Query, reply *conn.Reply) error {
	buf := bytes.NewBuffer(query.Data)
	dec := gob.NewDecoder(buf)
	var v helloRPC
	err := dec.Decode(&v)
	if err != nil {
		glog.Fatal("decode:", err)
	}

	addrs[v.Id] = v.Addr
	reply.Data = []byte(strconv.Itoa(int(cur_node.id)))

	fmt.Println("In Hello")
	return nil
}

func (w *Worker) JoinCluster(query *conn.Query, reply *conn.Reply) error {
	i, _ := strconv.Atoi(string(query.Data))
	id := uint64(i)
	cur_node.raft.ProposeConfChange(cur_node.ctx, raftpb.ConfChange{
		ID:      id,
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  id,
		Context: []byte(""),
	})
	return nil
}

func serveRequests(irwc io.ReadWriteCloser) {
	for {
		sc := &conn.ServerCodec{
			Rwc: irwc,
		}
		rpc.ServeRequest(sc)
	}
}

func runServer(address string) error {
	ln, err := net.Listen("tcp", address)
	if err != nil {
		glog.Fatalf("While running server: %v", err)
		return err
	}
	glog.WithField("address", ln.Addr()).Info("Worker listening")

	go func() {
		for {
			cxn, err := ln.Accept()
			if err != nil {
				glog.Fatalf("listen(%q): %s\n", address, err)
				return
			}
			glog.WithField("local", cxn.LocalAddr()).
				WithField("remote", cxn.RemoteAddr()).
				Debug("Worker accepted connection")
			go serveRequests(cxn)
		}
	}()
	return nil
}

func newNode(id uint64, addr string, peers []raft.Peer) *node {
	store := raft.NewMemoryStorage()
	n := &node{
		id:    id,
		addr:  addr,
		ctx:   context.TODO(),
		store: store,
		cfg: &raft.Config{
			ID:              id,
			ElectionTick:    5 * hb,
			HeartbeatTick:   hb,
			Storage:         store,
			MaxSizePerMsg:   math.MaxUint16,
			MaxInflightMsgs: 256,
		},
		pstore: make(map[string]string),
		ticker: time.Tick(time.Second),
		done:   make(chan struct{}),
	}

	n.raft = raft.StartNode(n.cfg, peers)
	return n
}

func (n *node) run() {
	for {
		select {
		case <-n.ticker:
			n.raft.Tick()
		case rd := <-n.raft.Ready():
			n.saveToStorage(rd.HardState, rd.Entries, rd.Snapshot)
			n.send(rd.Messages)
			if !raft.IsEmptySnap(rd.Snapshot) {
				n.processSnapshot(rd.Snapshot)
			}
			for _, entry := range rd.CommittedEntries {
				n.process(entry)
				if entry.Type == raftpb.EntryConfChange {
					var cc raftpb.ConfChange
					cc.Unmarshal(entry.Data)
					n.raft.ApplyConfChange(cc)
				}
			}
			n.raft.Advance()
		case <-n.done:
			return
		}
	}
}

func (n *node) saveToStorage(hardState raftpb.HardState, entries []raftpb.Entry, snapshot raftpb.Snapshot) {
	n.store.Append(entries)

	if !raft.IsEmptyHardState(hardState) {
		n.store.SetHardState(hardState)
	}

	if !raft.IsEmptySnap(snapshot) {
		n.store.ApplySnapshot(snapshot)
	}
}

func (n *node) send(messages []raftpb.Message) {
	for _, m := range messages {
		log.Println(raft.DescribeMessage(m, nil))

		// send message to other node
		//nodes[int(m.To)].receive(n.ctx, m)
		sendOverNetwork(n.ctx, m)
	}
}

func sendOverNetwork(ctx context.Context, message raftpb.Message) {
	pool, ok := pools[message.To]
	if !ok {
		connectWith(addrs[message.To])
		pool = pools[message.To]
	}
	addr := pool.Addr
	fmt.Println(addr)
	query := new(conn.Query)

	var network bytes.Buffer
	gob.Register(ctx)
	enc := gob.NewEncoder(&network)
	err := enc.Encode(raftRPC{ctx, message})
	if err != nil {
		glog.Fatalf("encode:", err)
	}

	query.Data = network.Bytes()
	reply := new(conn.Reply)
	if err := pool.Call("Worker.ReceiveOverNetwork", query, reply); err != nil {
		glog.WithField("call", "Worker.ReceiveOverNetwork").Fatal(err)
	}
	glog.WithField("reply_len", len(reply.Data)).WithField("addr", addr).
		Info("Got reply from server")

}

func (w *Worker) ReceiveOverNetwork(query *conn.Query, reply *conn.Reply) error {
	buf := bytes.NewBuffer(query.Data)
	dec := gob.NewDecoder(buf)
	gob.Register(context.Background())
	var v raftRPC
	err := dec.Decode(&v)
	if err != nil {
		glog.Fatal("decode:", err)
	}
	cur_node.receive(v.Ctx, v.Message)

	return nil
}

func (n *node) processSnapshot(snapshot raftpb.Snapshot) {
	panic(fmt.Sprintf("Applying snapshot on node %v is not implemented", n.id))
}

func (n *node) process(entry raftpb.Entry) {
	log.Printf("node %v: processing entry: %v\n", n.id, entry)
	if entry.Type == raftpb.EntryNormal && entry.Data != nil {
		parts := bytes.SplitN(entry.Data, []byte(":"), 2)
		n.pstore[string(parts[0])] = string(parts[1])
	}
}

func (n *node) receive(ctx context.Context, message raftpb.Message) {
	n.raft.Step(ctx, message)
}

func connectWith(addr string) uint64 {
	if len(addr) == 0 {
		return 0
	}
	pool := conn.NewPool(addr, 5)
	query := new(conn.Query)
	var network bytes.Buffer
	enc := gob.NewEncoder(&network)
	err := enc.Encode(helloRPC{cur_node.id, *workerPort})
	if err != nil {
		glog.Fatalf("encode:", err)
	}
	query.Data = network.Bytes()

	reply := new(conn.Reply)
	if err := pool.Call("Worker.Hello", query, reply); err != nil {
		glog.WithField("call", "Worker.Hello").Fatal(err)
	}
	i, _ := strconv.Atoi(string(reply.Data))
	glog.WithField("reply", i).WithField("addr", addr).
		Info("Got reply from server")

	pools[uint64(i)] = pool
	return uint64(i)
}

func proposeJoin(id uint64) {
	pool := pools[id]
	addr := pool.Addr
	fmt.Println(addr)
	query := new(conn.Query)
	query.Data = []byte(strconv.Itoa(int(cur_node.id)))
	reply := new(conn.Reply)
	if err := pool.Call("Worker.JoinCluster", query, reply); err != nil {
		glog.WithField("call", "Worker.JoinCluster").Fatal(err)
	}
	glog.WithField("reply_len", len(reply.Data)).WithField("addr", addr).
		Info("Got reply from server")
}

var (
	nodes      = make(map[int]*node)
	w          = new(Worker)
	workerPort = flag.String("workerport", ":12345",
		"Port used by worker for internal communication.")
	instanceIdx = flag.Uint64("idx", 1,
		"raft instance id")
	cluster  = flag.String("clusterIP", "", "IP of a node in cluster")
	cur_node *node
)

func main() {
	flag.Parse()
	cur_node = newNode(*instanceIdx, "", []raft.Peer{{ID: *instanceIdx}})
	if err := rpc.Register(w); err != nil {
		glog.Fatal(err)
	}
	if err := runServer(*workerPort); err != nil {
		glog.Fatal(err)
	}

	if *cluster != "" {
		i := connectWith(*cluster)
		go cur_node.run()
		proposeJoin(i)
	} else {
		go cur_node.run()
	}

	for cur_node.id == 1 && cur_node.raft.Status().Lead != 1 {
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("proposal by node ", cur_node.id)
	nodeID := strconv.Itoa(int(cur_node.id))
	cur_node.raft.Propose(cur_node.ctx, []byte("mykey"+nodeID+":myvalue"+nodeID))

	for cur_node.raft.Status().Lead != 1 {
		time.Sleep(100 * time.Millisecond)
	}

	count := 0
	for count != 3 {
		count = 0
		fmt.Printf("** Node %v **\n", cur_node.id)
		for k, v := range cur_node.pstore {
			fmt.Printf("%v = %v\n", k, v)
			count += 1
		}
		fmt.Printf("*************\n")
		time.Sleep(1000 * time.Millisecond)
	}

	time.Sleep(1000 * time.Millisecond)

}
