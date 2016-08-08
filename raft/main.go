package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"

	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
)

type node struct {
	cfg   *raft.Config
	ctx   context.Context
	data  map[string]string
	done  <-chan struct{}
	id    uint64
	raft  raft.Node
	store *raft.MemoryStorage
}

var (
	nodes = make(map[int]*node)
)

// HardState contains term, vote and commit.
// Snapshot contains data and snapshot metadata.
func (n *node) saveToStorage(hardState raftpb.HardState,
	entries []raftpb.Entry, snapshot raftpb.Snapshot) {

	n.store.Append(entries)
	if !raft.IsEmptyHardState(hardState) {
		n.store.SetHardState(hardState)
	}

	if !raft.IsEmptySnap(snapshot) {
		n.store.ApplySnapshot(snapshot)
	}
}

// receive a single message.
func (n *node) receive(ctx context.Context, message raftpb.Message) {
	n.raft.Step(ctx, message)
}

// send messages to peers.
func (n *node) send(messages []raftpb.Message) {
	for _, m := range messages {
		log.Println(raft.DescribeMessage(m, nil))

		nodes[int(m.To)].receive(context.TODO(), m)
	}
}

func (n *node) processSnapshot(snapshot raftpb.Snapshot) {
	log.Fatalf("Applying snapshot on node %v is not implemented", n.id)
}

func (n *node) process(entry raftpb.Entry) {
	log.Printf("node %v: processing entry: %v\n", n.id, entry)
	if entry.Type == raftpb.EntryNormal && entry.Data != nil {
		parts := bytes.SplitN(entry.Data, []byte(":"), 2)
		n.data[string(parts[0])] = string(parts[1])
	}
}

func (n *node) run() {
	for {
		select {
		case <-time.Tick(5 * time.Second):
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

func newNode(id uint64, peers []raft.Peer) *node {
	store := raft.NewMemoryStorage()
	n := &node{
		id:    id,
		store: store,
		cfg: &raft.Config{
			ID:              uint64(id),
			ElectionTick:    3,
			HeartbeatTick:   1,
			Storage:         store,
			MaxSizePerMsg:   4096,
			MaxInflightMsgs: 256,
		},
		data: make(map[string]string),
		ctx:  context.TODO(),
	}
	n.raft = raft.StartNode(n.cfg, peers)
	return n
}

func main() {
	// start a small cluster
	nodes[1] = newNode(1, []raft.Peer{{ID: 1}, {ID: 2}, {ID: 3}})
	nodes[1].raft.Campaign(nodes[1].ctx)
	go nodes[1].run()

	nodes[2] = newNode(2, []raft.Peer{{ID: 1}, {ID: 2}, {ID: 3}})
	go nodes[2].run()

	nodes[3] = newNode(3, []raft.Peer{})
	go nodes[3].run()
	nodes[2].raft.ProposeConfChange(nodes[2].ctx, raftpb.ConfChange{
		ID:      3,
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  3,
		Context: []byte(""),
	})

	// Wait for leader, is there a better way to do this
	for nodes[1].raft.Status().Lead != 1 {
		time.Sleep(100 * time.Millisecond)
	}

	nodes[1].raft.Propose(nodes[2].ctx, []byte("mykey1:myvalue1"))
	nodes[2].raft.Propose(nodes[2].ctx, []byte("mykey2:myvalue2"))
	nodes[3].raft.Propose(nodes[2].ctx, []byte("mykey3:myvalue3"))

	// Wait for proposed entry to be commited in cluster.
	// Apperently when should add an uniq id to the message and wait until it is
	// commited in the node.
	fmt.Printf("** Sleeping to visualize heartbeat between nodes **\n")
	time.Sleep(2000 * time.Millisecond)

	// Just check that data has been persited
	for i, node := range nodes {
		fmt.Printf("** Node %v **\n", i)
		for k, v := range node.data {
			fmt.Printf("%v = %v\n", k, v)
		}
		fmt.Printf("*************\n")
	}
}
