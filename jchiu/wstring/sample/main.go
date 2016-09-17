package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/jchiu0/experimental/wstring"
)

func main() {
	s := wstring.NewWString()
	s.Set([]byte("helloworld")) // This is a copy. No worries about double freeing.
	data := s.Get()
	fmt.Printf("[%s]\n", string(data))
	fmt.Printf("Length = %d\n", s.Size())

	// s is no longer used. We expect it to be destroyed when GC runs.
	// We want to make sure that std::string's destructor is run.
	// Make sure we see "Destroying wstring" before we see "Exiting".
	fmt.Println("GC start")
	runtime.GC()
	fmt.Println("GC end")
	time.Sleep(time.Second) // Give GC a bit of time.
	fmt.Println("Exiting")
}
