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
		for which := 0; which < 2; which++ {
			name := "Proto"
			if which == 1 {
				name = "Flatb"
			}

			b.Run(fmt.Sprintf("%s-%d", name, bm.k), func(b *testing.B) {
				uids := make([]uint64, bm.k)
				for i := 0; i < bm.k; i++ {
					uids[i] = uint64(rand.Int63())
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					var err error
					if which == 0 {
						err = ToAndFromProto(uids)
					} else {
						err = ToAndFromFlat(uids)
					}
					if err != nil {
						b.Error(err)
						b.Fail()
					}
				}
			})
		}
	}
}
