package flats

import (
	"fmt"
	"math/rand"
	"testing"
)

func BenchmarkToAndFromProto(b *testing.B) {
	benchmarks := []struct {
		k int
	}{
		{10},
		{100},
		{1000},
		{10000},
		{100000},
		{1000000},
	}

	for _, bm := range benchmarks {
		b.Run(fmt.Sprintf("%d", bm.k), func(b *testing.B) {
			uids := make([]uint64, bm.k)
			for i := 0; i < bm.k; i++ {
				uids[i] = uint64(rand.Int63())
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				if err := ToAndFromProto(uids); err != nil {
					b.Error(err)
					b.Fail()
				}
			}
		})
	}
}
