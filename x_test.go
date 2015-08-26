package x

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// Around 0.37 ns/op on my laptop.
// This is 25x faster than string comparisons, so iterating over and merging
// lists of ints would be a lot faster than doing the same for strings.
func BenchmarkInt64(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var m, n int64
	m = rand.Int63()
	n = rand.Int63()
	fmt.Printf("m=[%v] n=[%v]\n", m, n)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m == n
	}
}

// There's no difference between this and int64 benchmarks.
func BenchmarkInt32(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var m, n int32
	m = rand.Int31()
	n = rand.Int31()
	fmt.Printf("m=[%v] n=[%v]\n", m, n)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m == n
	}
}

// Around 8.67 ns/op on my laptop.
func BenchmarkString(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var m, n string
	l := rand.Intn(11) // Num permutations are now 5x +ve vals for int64
	m = UniqueString(l)
	n = UniqueString(l)
	fmt.Printf("length is: %v m=[%v] n=[%v]\n", l, m, n)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m == n
	}
}
