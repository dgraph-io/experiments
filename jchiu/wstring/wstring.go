package wstring

// #include <stdint.h>
// #include <stdlib.h>
// #include "wstring.h"
import "C"

import (
	"runtime"
)

type WString struct {
	c *C.wstring_t
}

func NewWString() *WString {
	s := &WString{C.wstring_new()}
	runtime.SetFinalizer(s, destroyWString)
	return s
}

func destroyWString(s *WString) {
	C.wstring_destroy(s.c)
}

// Get returns a slice of the string data. Note that there is no copying here.
// Do NOT mess with return value. It is a const char* in C++ and is readonly.
func (s *WString) Get() []byte {
	var l C.size_t
	data := C.wstring_get(s.c, &l)
	return charToByte(data, l)
}

// Set assigns the C++ string to the given data. There is copying here.
func (s *WString) Set(data []byte) {
	C.wstring_set(s.c, byteToChar(data), C.size_t(len(data)))
}

func (s *WString) Size() int {
	return int(C.wstring_len(s.c))
}
