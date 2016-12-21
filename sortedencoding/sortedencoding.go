package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"
)

func main() {
	a := 2<<24 + 10
	b := -2<<24 - 1
	arr := []int{a, b, 1, 2, 3, 4, -1, -2, -3, 0, 234, 10000, 123, -1543}
	sarr := make([]string, 0)
	for _, it := range arr {
		buf := new(bytes.Buffer)
		//Encode using bigendian.
		if it < 0 {
			buf.WriteByte(0)
		} else {
			buf.WriteByte(1)
		}
		err := binary.Write(buf, binary.BigEndian, int32(it))
		b := buf.Bytes()
		// Filp the last but so we can preserve ordering.
		if err != nil {
			fmt.Println(err)
		}
		sarr = append(sarr, string(b))
	}
	sort.Sort(sort.IntSlice(arr))
	sort.Sort(sort.StringSlice(sarr))
	fmt.Println(arr)
	for _, it := range sarr {
		var pi int32
		fmt.Printf("%v ", []byte(it))
		itOrig := []byte(it[1:])
		// Flip it back before decoding
		buf := bytes.NewReader(itOrig)
		// Decode to get original value
		err := binary.Read(buf, binary.BigEndian, &pi)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n", pi)
	}
}
