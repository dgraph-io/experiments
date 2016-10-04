package chanqueue

import (
	"sync"
	"testing"

	"github.com/textnode/gringo"
)

func BenchmarkChan(b *testing.B) {
	c := make(chan struct{}, 10)
	b.StartTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			c <- struct{}{}
		}
		close(c)
	}()
	for _ = range c {
	}
}

func BenchmarkQueue(b *testing.B) {
	q := NewQueue()
	var wg sync.WaitGroup
	wg.Add(1)
	b.StartTimer()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			q.Push()
		}
	}()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
	wg.Wait()
}

func BenchmarkCQueue(b *testing.B) {
	q := NewCQueue()
	var wg sync.WaitGroup
	wg.Add(1)
	b.StartTimer()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			q.Push()
		}
	}()
	for i := 0; i < b.N; i++ {
		q.Pop()
	}
	wg.Wait()
}

var payload = *gringo.NewPayload(1)

func BenchmarkGringo(b *testing.B) {
	q := gringo.NewGringo()
	var wg sync.WaitGroup
	wg.Add(1)
	b.StartTimer()
	go func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			q.Write(payload)
		}
	}()
	for i := 0; i < b.N; i++ {
		q.Read()
	}
	wg.Wait()
}
