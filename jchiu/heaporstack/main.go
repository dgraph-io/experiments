package main

import (
	"fmt"
)

type S struct{}

func getNew() *S {
	return &S{}
}

func identity(x *S) *S { return x }

func main() {
	a := getNew()
	identity(a)

	var b S
	identity(&b)

	c := new(S)
	identity(c)

	d := new(S)
	fmt.Println(d)

	var e S
	fmt.Println(&e)
}
