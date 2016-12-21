package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"sort"
)

func main() {
	arr := []int{1, 2, 3, 4, -1, -2, -3, 0, 234, 10000, 123, -1543}
	sarr := make([]string, 0)
	for _, it := range arr {
		buf := new(bytes.Buffer)
		//Encode using bigendian.
		err := binary.Write(buf, binary.BigEndian, int32(it))
		b := buf.Bytes()
		// Filp the last but so we can preserve ordering.
		b[0] = ^b[0]
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
		itOrig := []byte(it)
		// Flip it back before decoding
		itOrig[0] = ^itOrig[0]
		buf := bytes.NewReader(itOrig)
		// Decode to get original value
		err := binary.Read(buf, binary.BigEndian, &pi)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%v\n", pi)
	}
}
