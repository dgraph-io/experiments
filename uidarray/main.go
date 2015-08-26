// This is how CapnProto works:
// brew install capnp
// go get -v github.com/glycerine/go-capnproto
// copy over go.capnnp
//
// go get -v -t github.com/glycerine/bambam
// For bambam
// # ignore the initial compile error about 'undefined: LASTGITCOMMITHASH'. `make` will fix that.
// $ cd $GOPATH/src/github.com/glycerine/bambam
// $ make  # runs tests, build if all successful
// $ go install
// bambam -o . -p main main.go
// This generates schema.capnp
//
// Now compile all *.capnp with:
// capnp compile -ogo *.capnp
// No do a go build .

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
)

type UidArray struct {
	Uids []uint64 `capid:"0"`
}

func main() {
	u := UidArray{}
	for i := 0; i < 10; i++ {
		u.Uids = append(u.Uids, uint64(rand.Int63()))
	}
	fmt.Println(u)

	var b bytes.Buffer
	foo := bufio.NewWriter(&b)
	if err := u.Save(foo); err != nil {
		log.Fatalf("While writing to buffer: %v", err)
		return
	}
	foo.Flush()

	var ud UidArray
	ud.Load(&b)
	fmt.Println("Final")
	fmt.Println(ud)
}
