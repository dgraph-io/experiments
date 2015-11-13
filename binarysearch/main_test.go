package main

import (
	"math/rand"
	"testing"
)

func benchRecr(b *testing.B, sz int) {
	ar := make([]int, sz)
	for i := 0; i < len(ar); i++ {
		ar[i] = 3 * i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := rand.Intn(3 * sz)
		findSmallerOrEqualsRecr(ar, val, 0, sz-1)
	}
}

func BenchmarkRec_100(b *testing.B)   { benchRecr(b, 100) }
func BenchmarkRec_1000(b *testing.B)  { benchRecr(b, 1000) }
func BenchmarkRec_10000(b *testing.B) { benchRecr(b, 10000) }

func benchIter(b *testing.B, sz int) {
	ar := make([]int, sz)
	for i := 0; i < len(ar); i++ {
		ar[i] = 3 * i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := rand.Intn(3 * sz)
		findSmallerOrEqualsIter(ar, val)
	}
}

func BenchmarkIter_100(b *testing.B)   { benchIter(b, 100) }
func BenchmarkIter_1000(b *testing.B)  { benchIter(b, 1000) }
func BenchmarkIter_10000(b *testing.B) { benchIter(b, 10000) }

func benchLinear(b *testing.B, sz int) {
	ar := make([]int, sz)
	for i := 0; i < len(ar); i++ {
		ar[i] = 3 * i
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		val := rand.Intn(3 * sz)
		findSmallerOrEqualsLinear(ar, val)
	}
}

func BenchmarkLinear_100(b *testing.B)   { benchLinear(b, 100) }
func BenchmarkLinear_1000(b *testing.B)  { benchLinear(b, 1000) }
func BenchmarkLinear_10000(b *testing.B) { benchLinear(b, 10000) }
