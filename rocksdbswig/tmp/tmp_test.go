package tmp

import (
	"fmt"
	"testing"
)

func Test1(t *testing.T) {
	var s string
	SetMyString(&s)
	fmt.Println(s)
}
