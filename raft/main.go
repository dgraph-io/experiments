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

	if !raft.IsEmptySnap(snapshot) {
		fmt.Printf("saveToStorage snapshot: %v\n", snapshot.String())
		le, err := n.store.LastIndex()
		if err != nil {
			log.Fatalf("While retrieving last index: %v\n", err)
		}
		te, err := n.store.Term(le)
		if err != nil {
			log.Fatalf("While retrieving term: %v\n", err)
		}
		fmt.Printf("%d node Term for le: %v is %v\n", n.id, le, te)
		if snapshot.Metadata.Index <= le {
			fmt.Printf("%d node ignoring snapshot. Last index: %v\n", n.id, le)
			return
		}

		if err := n.store.ApplySnapshot(snapshot); err != nil {
			log.Fatalf("Applying snapshot: %v", err)
		}
	}

	if !raft.IsEmptyHardState(hardState) {
		n.store.SetHardState(hardState)
	}
	n.store.Append(entries)
}

// receive a single message.
func (n *node) receive(ctx context.Context, message raftpb.Message) {
	n.raft.Step(ctx, message)
}

// send messages to peers.
func (n *node) send(messages []raftpb.Message) {
	for _, m := range messages {
		log.Println("SEND: ", raft.DescribeMessage(m, nil))

		nodes[int(m.To)].receive(context.TODO(), m)
		if m.Type == raftpb.MsgSnap {
			// 	n.raft.ReportSnapshot(m.To, raft.SnapshotFinish)
		}
	}
}

// processSnapshot applies the snapshot to state machine.
func (n *node) applyToStateMachine(snapshot raftpb.Snapshot) {
	lead := int(n.raft.Status().Lead)
	for k, v := range nodes[lead].data {
		n.data[k] = v
	}
}

func (n *node) process(entry raftpb.Entry) {
	fmt.Printf("node %v: processing entry", n.id)
	if entry.Data == nil {
		return
	}
	if entry.Type == raftpb.EntryConfChange {
		fmt.Printf("Configuration change\n")
		var cc raftpb.ConfChange
		cc.Unmarshal(entry.Data)
		n.raft.ApplyConfChange(cc)
		return
	}

	if entry.Type == raftpb.EntryNormal {
		parts := bytes.SplitN(entry.Data, []byte(":"), 2)
		k := string(parts[0])
		v := string(parts[1])
		n.data[k] = v
		fmt.Printf(" Key: %v Val: %v\n", k, v)
	}
}

func (n *node) run() {
	for {
		select {
		case <-time.Tick(time.Second):
			n.raft.Tick()
		case rd := <-n.raft.Ready():
			n.saveToStorage(rd.HardState, rd.Entries, rd.Snapshot)
			n.send(rd.Messages)
			if !raft.IsEmptySnap(rd.Snapshot) {
				fmt.Println("Applying snapshot to state machine")
				n.applyToStateMachine(rd.Snapshot)
			}
			if len(rd.CommittedEntries) > 0 {
				fmt.Printf("Node: %v. Got %d committed entries\n", n.id, len(rd.CommittedEntries))
			}
			for _, entry := range rd.CommittedEntries {
				n.process(entry)
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
	nodes[1] = newNode(1, []raft.Peer{{ID: 1}, {ID: 2}})
	go nodes[1].run()

	nodes[2] = newNode(2, []raft.Peer{{ID: 1}, {ID: 2}})
	go nodes[2].run()

	nodes[1].raft.Campaign(nodes[1].ctx)

	time.Sleep(10 * time.Second)

	fmt.Println("----------------- Adding NODE 3")
	nodes[3] = newNode(3, []raft.Peer{})
	go nodes[3].run()
	nodes[2].raft.ProposeConfChange(nodes[2].ctx, raftpb.ConfChange{
		ID:      3,
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  3,
		Context: []byte(""),
	})
	time.Sleep(10 * time.Second)
	fmt.Println("------------------ Proposing values")

	// Wait for leader, is there a better way to do this
	/*
		for nodes[1].raft.Status().Lead != 1 {
			fmt.Println("Waiting for 1 to become leader")
			time.Sleep(100 * time.Millisecond)
		}
	*/

	nodes[1].raft.Propose(nodes[2].ctx, []byte("mykey1:myvalue1"))
	nodes[2].raft.Propose(nodes[2].ctx, []byte("mykey2:myvalue2"))
	nodes[3].raft.Propose(nodes[2].ctx, []byte("mykey3:myvalue3"))

	// Wait for proposed entry to be commited in cluster.
	// Apperently when should add an uniq id to the message and wait until it is
	// commited in the node.
	fmt.Printf("** Sleeping to visualize heartbeat between nodes **\n")
	time.Sleep(5 * time.Second)

	leader := int(nodes[1].raft.Status().Lead)
	fmt.Printf("=========== Taking snapshot of leader: %v\n", leader)
	le, err := nodes[leader].store.LastIndex()
	if err != nil {
		log.Fatalf("node leader: %v has error getting last index: err: %v\n.", le, err)
	}
	_, err = nodes[leader].store.CreateSnapshot(le-1, nil, []byte(""))
	if err != nil {
		log.Fatalf("node leader: %v has error taking snapshot: %v\n.", le, err)
	}
	if err := nodes[leader].store.Compact(le - 1); err != nil {
		log.Fatalf("node leader: %v has error while compaction: %v\n", le, err)
	}
	fmt.Println("===================== Snapshot taken")

	// Just check that data has been persited
	for i, node := range nodes {
		fmt.Printf("** Node %v **\n", i)
		for k, v := range node.data {
			fmt.Printf("%v = %v\n", k, v)
		}
		fmt.Printf("*************\n")
	}

	fmt.Println("Adding a new 4th node")
	nodes[4] = newNode(4, []raft.Peer{{ID: 1}})
	go nodes[4].run()
	nodes[1].raft.ProposeConfChange(nodes[1].ctx, raftpb.ConfChange{
		ID:      4,
		Type:    raftpb.ConfChangeAddNode,
		NodeID:  4,
		Context: []byte(""),
	})
	fmt.Printf("** Sleeping to visualize heartbeat between nodes **\n")
	time.Sleep(5 * time.Second)

	// Just check that data has been propagated to 4.
	for i, node := range nodes {
		fmt.Printf("** Node %v **\n", i)
		for k, v := range node.data {
			fmt.Printf("%v = %v\n", k, v)
		}
		fmt.Printf("*************\n")
	}
}
