package chanqueue

import (
	"testing"
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
	b.StartTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			q.Push()
		}
		q.Done()
	}()
	for !q.IsDone() {
		q.Pop()
	}
}

func BenchmarkCQueue(b *testing.B) {
	q := NewCQueue()
	b.StartTimer()
	go func() {
		for i := 0; i < b.N; i++ {
			q.Push()
		}
		q.Done()
	}()
	for !q.IsDone() {
		q.Pop()
	}
}
