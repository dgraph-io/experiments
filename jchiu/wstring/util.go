package wstring

import "C"

import (
	"reflect"
	"unsafe"
)

// byteToChar returns *C.char from byte slice.
func byteToChar(b []byte) *C.char {
	var c *C.char
	if len(b) > 0 {
		c = (*C.char)(unsafe.Pointer(&b[0]))
	}
	return c
}

// charToByte converts a *C.char to a byte slice.
func charToByte(data *C.char, l C.size_t) []byte {
	var value []byte
	sH := (*reflect.SliceHeader)(unsafe.Pointer(&value))
	sH.Cap, sH.Len, sH.Data = int(l), int(l), uintptr(unsafe.Pointer(data))
	return value
}
