package benchhash

import (
	"sync"
)

type GoMap struct {
	sync.RWMutex
	m map[uint32]uint32
}

func (s *GoMap) Get(key uint32) (uint32, bool) {
	s.RLock()
	defer s.RUnlock()
	val, found := s.m[key]
	return val, found
}

func (s *GoMap) Put(key, val uint32) {
	s.Lock()
	defer s.Unlock()
	s.m[key] = val
}

func NewGoMap() HashMap {
	return &GoMap{
		m: make(map[uint32]uint32),
	}
}

type ShardedGoMap struct {
	numShards int
	m         []HashMap
}

func (s *ShardedGoMap) Get(key uint32) (uint32, bool) {
	shard := key % uint32(s.numShards)
	return s.m[shard].Get(key)
}

func (s *ShardedGoMap) Put(key, val uint32) {
	shard := key % uint32(s.numShards)
	s.m[shard].Put(key, val)
}

func NewShardedGoMap(numShards int) HashMap {
	r := &ShardedGoMap{
		numShards: numShards,
		m:         make([]HashMap, numShards),
	}
	for i := 0; i < numShards; i++ {
		r.m[i] = NewGoMap()
	}
	return r
}

func NewShardedGoMap4() HashMap  { return NewShardedGoMap(4) }
func NewShardedGoMap8() HashMap  { return NewShardedGoMap(8) }
func NewShardedGoMap16() HashMap { return NewShardedGoMap(16) }
func NewShardedGoMap32() HashMap { return NewShardedGoMap(32) }
