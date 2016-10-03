package chanqueue

import (
	"sync"
)

type Queue struct {
	sync.RWMutex
	data []struct{} // Size should remain constant throughout.
	idx  int        // Where to pop.
	done bool
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]struct{}, 0, 1000),
	}
}

// Done marks queue as done.
func (q *Queue) Done() {
	q.Lock()
	defer q.Unlock()
	q.done = true
}

// Done checks if queue is done.
func (q *Queue) IsDone() bool {
	q.RLock()
	defer q.RUnlock()
	return q.done && q.idx == len(q.data)
}

func (q *Queue) Push() {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, struct{}{})
}

func (q *Queue) IsEmpty() bool {
	q.RLock()
	defer q.RUnlock()
	return q.idx == len(q.data)
}

// Pop returns an item.
func (q *Queue) Pop() {
	q.Lock()
	defer q.Unlock()
	if q.idx >= len(q.data) {
		return
	}
	q.idx++
}
