package chanqueue

import (
	"sync"
)

type CQueue struct {
	sync.RWMutex
	data []struct{} // Size should remain constant throughout.
	out  int        // Where to pop.
	in   int        // Where to push.
	done bool
}

func NewCQueue() *CQueue {
	return &CQueue{
		data: make([]struct{}, 1000),
	}
}

// Done marks queue as done.
func (q *CQueue) Done() {
	q.Lock()
	defer q.Unlock()
	q.done = true
}

// Done checks if queue is done.
func (q *CQueue) IsDone() bool {
	q.RLock()
	defer q.RUnlock()
	return q.done && (q.in == q.out)
}

func (q *CQueue) tryPush() bool {
	q.Lock()
	defer q.Unlock()
	newIn := (q.in + 1) % len(q.data)
	if newIn == q.out {
		return false
	}
	// q.data[q.in] = new element.
	q.in = newIn
	return true
}

func (q *CQueue) Push() {
	for !q.tryPush() {
	}
}

func (q *CQueue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()
	return q.in == q.out
}

func (q *CQueue) tryPop() bool {
	q.Lock()
	defer q.Unlock()
	if q.in == q.out {
		// Queue is empty.
		return false
	}
	// Element to return is q.data[q.out].
	q.out = (q.out + 1) % len(q.data)
	return true
}

// Pop returns an item.
func (q *CQueue) Pop() {
	for !q.tryPop() {
	}
}
