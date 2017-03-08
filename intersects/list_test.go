package intersect

import (
	"fmt"
	"math"
	"math/rand"
	"sort"
	"testing"
	"time"

	"github.com/dgraph-io/dgraph/task"
	"github.com/stretchr/testify/require"
)

func TestUIDListIntersect2(t *testing.T) {
	u := []uint64{1, 2, 3}
	v := []uint64{1, 2, 3, 4, 5}
	res := make([]uint64, 0, 3)
	BinIntersect(u, v, &res)
	require.Equal(t, []uint64{1, 2, 3}, res)
	require.Equal(t, []uint64{1, 2, 3, 4, 5}, v)
}

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
	a := createArray(100, 1000)
	da := encodeDelta(a, 3)
	b := createArray(200, 1000)
	db := encodeDelta(b, 3)
	fmt.Println("a=", da.Buckets, da.Uids)
	fmt.Println("b=", db.Buckets, db.Uids)

	fmt.Printf("a=%v\nb=%v\n", a, b)
	final := make([]uint64, 0, 100)
	twoLevelLinear(da, db, &final)

	exp := make([]uint64, 0, 100)
	mergeIntersect(a, b, &exp)
	fmt.Printf("exp=%v\n", exp)
	require.Equal(t, exp, final)

	final = final[:0]
	twoLevelBinary(da, db, &final)
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

	//res3 := make([]uint64, 0, 100)
	//binIterative(a, b, &res3)

	exp := intersect(a, b)

	require.Equal(t, exp, res1, "merge not working")
	require.Equal(t, exp, res2, "binIntersect not working")
	//require.Equal(t, exp, res3, "binIterative not working")
}

func BenchmarkMarshal(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	rs := []int{1, 10, 50, 100, 500, 1000, 10000, 100000, 1000000}
	for _, r := range rs {
		u := createArray(r, math.MaxInt32)

		var dsize, fsize int
		b.Run(fmt.Sprintf("delta-%d", r),
			func(b *testing.B) {
				var total time.Duration
				for i := 0; i < b.N; i++ {
					d := encodeDelta(u, 32)
					data, err := d.Marshal()
					if err != nil {
						b.Fatalf("Error: %v", err)
					}
					// Assuming 128 MB per sec. => 128 bytes per micro.
					dur := time.Duration(len(data) / 128)
					total += dur
					time.Sleep(dur * time.Microsecond)
					if i == 0 {
						dsize = len(data)
					}

					var out DeltaList
					if err := out.Unmarshal(data); err != nil {
						b.Fatalf("Error: %v", err)
					}
				}
				// fmt.Printf("Total sleep: %v\n", total)
			})

		b.Run(fmt.Sprintf("fixed-%d", r),
			func(b *testing.B) {
				var total time.Duration
				for i := 0; i < b.N; i++ {
					d := encodeFixed(u)
					data, err := d.Marshal()
					if err != nil {
						b.Fatalf("Error: %v", err)
					}
					dur := time.Duration(len(data) / 128)
					total += dur
					// Assuming 128 MB per sec. => 128 bytes per micro.
					time.Sleep(dur * time.Microsecond)

					if i == 0 {
						fsize = len(data)
					}
					var out FixedList
					if err := out.Unmarshal(data); err != nil {
						b.Fatalf("Error: %v", err)
					}
				}
				// fmt.Printf("Total fixed sleep: %v\n", total)
			})
		fmt.Printf("SIZE Delta: %v Fixed: %v\n", dsize, fsize)
		fmt.Println()
	}
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
			d1 := encodeDelta(u1, 32)
			u2 := createArray(sz2, limit)
			d2 := encodeDelta(u2, 32)
			result := make([]uint64, 0, sz)

			u := &task.List{u1}
			v := &task.List{u2}
			ucopy := make([]uint64, len(u1), len(u1))
			copy(ucopy, u1)

			b.Run(fmt.Sprintf(":Cur:size=%d:overlap=%.2f:ratio=%d:", sz, overlap, r),
				func(b *testing.B) {
					for k := 0; k < b.N; k++ {
						//u.Uids = u.Uids[:sz]
						//copy(u.Uids, ucopy)
						IntersectWith(u, v)
					}
				})

			b.Run(fmt.Sprintf(":Bin:size=%d:overlap=%.2f:ratio=%d:", sz, overlap, r),
				func(b *testing.B) {
					for k := 0; k < b.N; k++ {
						BinIntersect(u2, u1, &result)
						result = result[:0]
					}
				})
			b.Run(fmt.Sprintf(":Mer:size=%d:overlap=%.2f:ratio=%d", sz, overlap, r),
				func(b *testing.B) {
					for k := 0; k < b.N; k++ {
						mergeIntersect(u1, u2, &result)
						result = result[:0]
					}
				})

			b.Run(fmt.Sprintf(":Two:size=%d:overlap=%.2f:ratio=%d", sz, overlap, r),
				func(b *testing.B) {
					var f func(a, b *DeltaList, final *[]uint64)
					if r < 500 {
						f = twoLevelLinear
					} else {
						f = twoLevelBinary
					}
					for k := 0; k < b.N; k++ {
						f(d1, d2, &result)
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
