package codec

import (
	"fmt"
	"log"

	"github.com/dgraph-io/experiments/grpc/fb"
)

type Buffer struct{}

func (cb *Buffer) Marshal(v interface{}) ([]byte, error) {
	fmt.Println("Marshal")
	p, ok := v.(*fb.Payload)
	if !ok {
		log.Fatal("Invalid type of struct")
	}
	return p.Data, nil
}

func (cb *Buffer) Unmarshal(data []byte, v interface{}) error {
	fmt.Println("Unmarshal")
	p, ok := v.(*fb.Payload)
	if !ok {
		log.Fatal("Invalid type of struct")
	}
	p.Data = data
	return nil
}

func (cb *Buffer) String() string {
	fmt.Println("String")
	return "codec.Buffer"
}
