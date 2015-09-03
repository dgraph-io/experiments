package vrpc

import (
	"flag"

	"math/rand"

	"github.com/Sirupsen/logrus"
)

var port = flag.String("port", "12345", "Port to bind on")

type Reply struct {
	Req   interface{}
	Hello string
}

type PostingList struct {
	Uids []uint64
}

var log = logrus.WithField("pkg", "vrpc")

func NewRequest(sz int) *PostingList {
	req := new(PostingList)
	req.Uids = make([]uint64, sz)
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
