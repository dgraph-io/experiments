package main

import (
	"fmt"
	"math/rand"
	"time"

	"golang.org/x/net/context"

	"github.com/Sirupsen/logrus"
	"github.com/coreos/etcd/raft"
	"github.com/coreos/etcd/raft/raftpb"
)

func save(rd raft.Ready, st *raft.MemoryStorage) error {
	if !raft.IsEmptyHardState(rd.HardState) {
		if err := st.SetHardState(rd.HardState); err != nil {
			return err
		}
	}

	if len(rd.Entries) > 0 {
		if err := st.Append(rd.Entries); err != nil {
			return err
		}
	}

	if !raft.IsEmptySnap(rd.Snapshot) {
		if err := st.ApplySnapshot(rd.Snapshot); err != nil {
			return err
		}
	}
	return nil
}

func startNode(id int, chans []chan raftpb.Message) {
	l := log.WithField("id", id)
	storage := raft.NewMemoryStorage()
	c := &raft.Config{
		ID:              uint64(id),
		ElectionTick:    10,
		HeartbeatTick:   1,
		Storage:         storage,
		MaxSizePerMsg:   4096,
		MaxInflightMsgs: 256,
	}
	var peers []raft.Peer
	for i := 1; i <= 3; i++ {
		if id == i {
			continue
		}
		peer := raft.Peer{ID: uint64(i)}
		peers = append(peers, peer)
	}
	l.WithField("peers", peers).Debug("Peers")

	n := raft.StartNode(c, peers)
	tick := time.Tick(3 * time.Second)
	for count := 0; ; count++ {
		l.Debug("Waiting for something to happen")
		select {
		case <-tick:
			l.Debug("Got a tick")
			n.Tick()
			{
				r := 1 + rand.Intn(3)
				l.WithField("r", r).Debugf("Got rand value")
				if id == r {
					// I should send some data.
					data := fmt.Sprintf("This is me: %v at count %v", id, count)
					l.WithField("data", data).Debug("Proposing some data")
					ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
					go n.Propose(ctx, []byte(data))
				}
			}

		case rd := <-n.Ready():
			l.Debug("Got ready")
			// Save initial state.
			if err := save(rd, storage); err != nil {
				log.WithField("error", err).Error("While saving")
				return
			}
			// Send all messages to other nodes.
			for i := 1; i <= 3; i++ {
				if id == i {
					continue
				}
				for _, msg := range rd.Messages {
					chans[i-1] <- msg
				}
			}
			// Apply Snapshot (already done) and CommittedEntries
			storage.Append(rd.CommittedEntries) // No config change in test.
			n.Advance()

		case msg := <-chans[id-1]:
			l.WithField("msg", msg.String()).Debug("GOT MESSAGE")
			ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
			n.Step(ctx, msg)
		}
	}
}

var log = logrus.WithField("package", "goraft")

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	chans := make([]chan raftpb.Message, 3)
	for i := range chans {
		chans[i] = make(chan raftpb.Message, 100)
	}
	go startNode(1, chans)
	go startNode(2, chans)
	go startNode(3, chans)

	tick := time.Tick(60 * time.Second)
	<-tick
	log.Debug("DONE")
}
