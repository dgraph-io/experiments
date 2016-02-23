package main

import (
	"encoding/gob"
	"io"
)

type Query struct {
	Data []byte
}

func GobEncode(q *Query, w io.Writer) error {
	enc := gob.NewEncoder(w)
	return enc.Encode(*q)
}

func Encode(data []byte, w io.Writer) error {
	_, err := w.Write(data)
	return err
}
