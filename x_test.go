package x

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestMergeInt(t *testing.T) {
	l := SortedInt(10)
	if !sort.IsSorted(int64arr(l)) {
		t.Errorf("Not sorted: [%v]\n", l)
		t.FailNow()
	}
	for i := 0; i < 50000; i++ {
		l = mergeInt(l, rand.Int63())
		if !sort.IsSorted(int64arr(l)) {
			t.Errorf("Not sorted: [%v]\n", l)
			t.FailNow()
		}
	}
}

func TestMergeString(t *testing.T) {
	l := SortedString(10)
	for !sort.StringsAreSorted(l) {
		t.Errorf("Not sorted: [%v]\n", l)
	}
	for i := 0; i < 5000; i++ {
		s := UniqueString(rand.Intn(11))
		l = mergeString(l, s)
		if !sort.StringsAreSorted(l) {
			t.Error("Strings are not sorted")
			t.FailNow()
		}
	}
}

// Around 0.37 ns/op on my laptop.
// This is 25x faster than string comparisons, so iterating over and merging
// lists of ints would be a lot faster than doing the same for strings.
func BenchmarkInt64(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	var m, n int64
	m = rand.Int63()
	n = rand.Int63()
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
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m == n
	}
}

func BenchmarkMergeSortedInt(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sz := rand.Intn(50000) + 1000
	l := SortedInt(sz)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		l = mergeInt(l, rand.Int63())
	}
}

func BenchmarkMergeSortedString(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sz := rand.Intn(50000) + 1000
	l := SortedString(sz)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := UniqueString(rand.Intn(11))
		l = mergeString(l, s)
	}
}

func BenchmarkFindIndexInt(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sz := rand.Intn(50000) + 1000
	l := SortedInt(sz)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		findIndexInt(l, rand.Int63())
	}
}

func BenchmarkFindIndexString(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	sz := rand.Intn(50000) + 1000
	l := SortedString(sz)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		s := UniqueString(rand.Intn(11))
		findIndexString(l, s)
	}
}
