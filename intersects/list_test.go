package intersect

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func intersect(a, b []uint64) []uint64 {
	m := make(map[uint64]struct{})
	for _, i := range a {
		m[i] = struct{}{}
	}
	out := make([]uint64, 0, 100)
	for _, j := range b {
		if _, ok := m[j]; ok {
			out = append(out, j)
		}
	}
	return out
}

func TestIntersect(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a := make([]uint64, 100)
	b := make([]uint64, 100)
	sz := int64(1000)
	ma := make(map[uint64]struct{})
	mb := make(map[uint64]struct{})
	for i := range a {
		for {
			ei := uint64(rand.Int63n(sz))
			if _, ok := ma[ei]; !ok {
				a[i] = ei
				ma[ei] = struct{}{}
				break
			}
		}

		for {
			ei := uint64(rand.Int63n(sz))
			if _, ok := mb[ei]; !ok {
				b[i] = ei
				mb[ei] = struct{}{}
				break
			}
		}
	}

	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	sort.Slice(b, func(i, j int) bool {
		return b[i] < b[j]
	})
	fmt.Println(a)
	fmt.Println(b)

	res1 := merge(a, b)
	res2 := make([]uint64, 0, 100)
	binIntersect(a, b, &res2)
	fmt.Println(res2)

	exp := intersect(a, b)
	fmt.Println(exp)

	require.Equal(t, exp, res1, "merge not working")
	require.Equal(t, exp, res2, "binIntersect not working")
}
