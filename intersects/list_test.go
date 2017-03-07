package intersect

import (
	"fmt"
	"math"
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

func createArray(sz int, limit int64) []uint64 {
	a := make([]uint64, sz)
	ma := make(map[uint64]struct{})
	for i := 0; i < sz; i++ {
		for {
			ei := uint64(rand.Int63n(limit))
			if _, ok := ma[ei]; !ok {
				a[i] = ei
				ma[ei] = struct{}{}
				break
			}
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	return a
}

func TestSize(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	sz := 1
	for i := 0; i < 5; i++ {
		sz *= 10
		a := createArray(sz, math.MaxInt32)
		dl := encodeDelta(a, 32)
		dd, err := dl.Marshal()
		require.Nil(t, err)

		fl := encodeFixed(a)
		fd, err := fl.Marshal()
		require.Nil(t, err)
		fmt.Printf("Size=%d Size of delta: %v fixed: %v\n", sz, len(dd), len(fd))

		mi := func(d []uint64) uint64 {
			var dm uint64
			for _, e := range d {
				if dm < e {
					dm = e
				}
			}
			return dm
		}
		dm := mi(dl.Uids)
		fm := mi(fl.Uids)
		fmt.Printf("Max delta: %v. Max fixed: %v. Bits delta: %v. Bits fixed: %v\n",
			dm, fm, math.Log2(float64(dm)), math.Log2(float64(fm)))
		fmt.Println()
	}
}

func TestTwoLevelIntersect(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a := createArray(50, 100)
	da := encodeDelta(a, 3)
	b := createArray(80, 100)
	db := encodeDelta(b, 3)
	fmt.Println("a=", da.Buckets, da.Uids)
	fmt.Println(db.Buckets, db.Uids)

	fmt.Printf("a=%v\nb=%v\n", a, b)
	final := make([]uint64, 0, 100)
	mergeDeltaIntersect(da, db, &final)

	exp := make([]uint64, 0, 100)
	mergeIntersect(a, b, &exp)
	fmt.Printf("exp=%v\n", exp)
	require.Equal(t, exp, final)
}

func TestIntersect(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	a := createArray(100, 100)
	b := createArray(100, 100)

	res1 := make([]uint64, 0, 100)
	mergeIntersect(a, b, &res1)

	res2 := make([]uint64, 0, 100)
	binIntersect(a, b, &res2)

	res3 := make([]uint64, 0, 100)
	binIterative(a, b, &res3)

	exp := intersect(a, b)

	require.Equal(t, exp, res1, "merge not working")
	require.Equal(t, exp, res2, "binIntersect not working")
	//require.Equal(t, exp, res3, "binIterative not working")
}

func BenchmarkListIntersect(b *testing.B) {
	randomTests := func(sz int, overlap float64) {
		rs := []int{1, 10, 50, 100, 500, 1000, 10000, 100000, 1000000}

		for _, r := range rs {
			sz2 := sz * r
			if sz2 > 1000000 {
				break
			}
			limit := int64(float64(sz2) / overlap)

			u1 := createArray(sz, limit)
			u2 := createArray(sz2, limit)
			result := make([]uint64, 0, sz)

			b.Run(fmt.Sprintf(":Bin:size=%d:overlap=%.2f:ratio=%d:", sz, overlap, r),
				func(b *testing.B) {
					for k := 0; k < b.N; k++ {
						BinIntersect(u2, u1, &result)
						result = result[:0]
					}
				})
			//b.Run(fmt.Sprintf(":Itr:size=%d:ratio=%d:overlap=%.2f:", r, sz, overlap),
			//func(b *testing.B) {
			//for k := 0; k < b.N; k++ {
			//binIterative(u1, u2, &result)
			//result = result[:0]
			//}
			//})
			b.Run(fmt.Sprintf(":Mer:size=%d:overlap=%.2f:ratio=%d", sz, overlap, r),
				func(b *testing.B) {
					for k := 0; k < b.N; k++ {
						mergeIntersect(u1, u2, &result)
						result = result[:0]
					}
				})
			fmt.Println()
		}
	}
	//randomTests(10, 0.01)
	//randomTests(100, 0.01)
	//randomTests(1000, 0.01)
	//randomTests(10000, 0.01)

	// Overlap has no effect on Bin numbers.
	overlaps := []float64{0.00001, 0.8}
	for _, overlap := range overlaps {
		randomTests(10, overlap)
		randomTests(100, overlap)
		randomTests(1000, overlap)
		randomTests(10000, overlap)
		fmt.Println()
	}

	//randomTests(10, 0.4)
	//randomTests(100, 0.4)
	//randomTests(1000, 0.4)
	//randomTests(10000, 0.4)
	//fmt.Println()

	//randomTests(10, 0.8)
	//randomTests(100, 0.8)
	//randomTests(1000, 0.8)
	//randomTests(10000, 0.8)
	//fmt.Println()
}
