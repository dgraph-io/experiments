package main

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"

	"github.com/dgraph-io/experiments/zrpc/plist"
	C "github.com/glycerine/go-capnproto"
)

func write(sz int) (result bytes.Buffer, err error) {
	seg := C.NewBuffer(nil)
	pl := plist.NewRootPostingList(seg)

	l := seg.NewUInt64List(sz)
	for i := 0; i < sz; i++ {
		l.Set(i, uint64(rand.Int63()))
	}

	pl.SetIds(l)
	// pl.SetTitle("capnproto")

	r, err := seg.WriteTo(&result)
	if err != nil {
		return result, err
	}
	fmt.Println("Written:", r)
	return result, nil
}

func read(buf bytes.Buffer) error {
	seg, err := C.ReadFromStream(&buf, nil)
	if err != nil {
		log.Print("While decoding")
		return err
	}

	pl := plist.ReadRootPostingList(seg)
	ids := pl.Ids()
	title := pl.Title()
	fmt.Printf("Num ids: [%v] Title: [%v]\n", ids.Len(), title)
	return nil
}

func main() {
	sz := 1000

	buf, err := write(sz)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Buffer len:", buf.Len())
	if err := read(buf); err != nil {
		log.Fatal(err)
	}
}
