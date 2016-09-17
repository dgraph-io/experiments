package benchhash

import (
	"log"
	"math/rand"
	"sync"
	"testing"
)

type HashMap interface {
	Get(key uint32) (uint32, bool)
	Put(key, val uint32)
}

type KeyValPair struct {
	Key, Val uint32
}

func intArray(n int) []uint32 {
	a := make([]uint32, n)
	for i := 0; i < n; i++ {
		a[i] = rand.Uint32()
	}
	return a
}

func intPairArray(n int) []KeyValPair {
	a := make([]KeyValPair, n)
	for i := 0; i < n; i++ {
		a[i] = KeyValPair{rand.Uint32(), rand.Uint32()}
	}
	return a
}

func check(n, q int) {
	if (n % q) != 0 {
		log.Fatalf("%d not divisible by %d", n, q)
	}
}

func workRange(n, q, j int) (int, int) {
	return (n / q) * j, (n / q) * (j + 1)
}

// MultiRead gets n items using q Go routines. Do not use a channel / locking queue.
func MultiRead(n, q int, newFunc func() HashMap, b *testing.B) {
	check(n, q)
	work := intArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		h := newFunc()
		var wg sync.WaitGroup
		for j := 0; j < q; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				start, end := workRange(n, q, j)
				for k := start; k < end; k++ {
					h.Get(work[k])
				}
			}(j)
		}
		wg.Wait()
	}
}

// MultiWrite writes n items using q Go routines.
func MultiWrite(n, q int, newFunc func() HashMap, b *testing.B) {
	check(n, q)
	work := intPairArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		h := newFunc()
		var wg sync.WaitGroup
		for j := 0; j < q; j++ {
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				start, end := workRange(n, q, j)
				for k := start; k < end; k++ {
					h.Put(work[k].Key, work[k].Val)
				}
			}(j)
		}
		wg.Wait()
	}
}

// ReadWrite does read and write in parallel.
// qRead is num goroutines for reading.
// qWrite is num goroutines for writing.
// Assume n divisible by (qRead + qWrite).
func ReadWrite(n, qRead, qWrite int, newFunc func() HashMap, b *testing.B) {
	q := qRead + qWrite
	check(n, q)
	work := intPairArray(n)
	b.StartTimer()
	for i := 0; i < b.N; i++ { // N reps.
		h := newFunc()
		var wg sync.WaitGroup
		for j := 0; j < qRead; j++ { // Read goroutines.
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				start, end := workRange(n, q, j)
				for k := start; k < end; k++ {
					h.Get(work[k].Key)
				}
			}(j)
		}

		for j := qRead; j < q; j++ { // Write goroutines.
			wg.Add(1)
			go func(j int) {
				defer wg.Done()
				start, end := workRange(n, q, j)
				for k := start; k < end; k++ {
					h.Put(work[k].Key, work[k].Val)
				}
			}(j)
		}
		wg.Wait()
	}
}
