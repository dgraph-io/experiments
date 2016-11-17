package flats

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkToAndFrom(b *testing.B) {
	benchmarks := []struct {
		k int
	}{
		{10},
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
		{10000000},
	}

	for _, bm := range benchmarks {
		for which := 0; which < 3; which++ {
			var name string
			if which == 0 {
				fmt.Println()
				name = "Flatb"
			} else if which == 1 {
				name = "Fixed"
			} else if which == 2 {
				name = "Proto"
			}

			b.Run(fmt.Sprintf("%s-%d", name, bm.k), func(b *testing.B) {
				uids := make([]uint64, bm.k)
				for i := 0; i < bm.k; i++ {
					uids[i] = uint64(rand.Int63())
				}

				var max, sz int
				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var err error
					if which == 0 {
						err, sz = ToAndFromFlat(uids)
					} else if which == 1 {
						err, sz = ToAndFromProtoAlt(uids)
					} else if which == 2 {
						err, sz = ToAndFromProto(uids)
					}
					if err != nil {
						b.Error(err)
						b.Fail()
					}
					if max < sz {
						max = sz
					}
					// runtime.GC() -- Actually makes FB looks worse.
				}
			})
		}
	}
}

func TestToAndFrom(t *testing.T) {
	ar := []int{10, 100, 1000, 10000, 100000, 1000000, 10000000}
	for _, k := range ar {
		fmt.Println()
		uids := make([]uint64, k)
		for i := 0; i < k; i++ {
			uids[i] = uint64(rand.Int63())
		}
		var err error
		var sz int
		err, sz = ToAndFromFlat(uids)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		fmt.Printf("Flatb k:%d sz:%d\n", k, sz)

		err, sz = ToAndFromProtoAlt(uids)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		fmt.Printf("Fixed k:%d sz:%d\n", k, sz)

		err, sz = ToAndFromProto(uids)
		if err != nil {
			t.Error(err)
			t.Fail()
		}
		fmt.Printf("Proto k:%d sz:%d\n", k, sz)
	}
}
