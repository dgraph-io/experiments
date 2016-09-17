package benchhash

import (
	"github.com/zond/gotomic"
)

type GotomicMap struct {
	h *gotomic.Hash
}

func (s GotomicMap) Get(key uint32) (uint32, bool) {
	val, found := s.h.Get(gotomic.IntKey(key))
	if val == nil {
		return 0, false
	}
	return uint32(val.(gotomic.IntKey)), found
}

func (s GotomicMap) Put(key, val uint32) {
	s.h.Put(gotomic.IntKey(key), gotomic.IntKey(val))
}

func NewGotomicMap() HashMap {
	return GotomicMap{gotomic.NewHash()}
}
