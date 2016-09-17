package benchhash2

import (
	"math/rand"
	"testing"
)

type experiment struct {
	label   string
	newFunc func() HashMap
}

var (
	experiments = []experiment{
		experiment{"GoMap", NewGoMap},
		experiment{"GotomicMap", NewGotomicMap},
		experiment{"ShardedGoMap8", NewShardedGoMap8},
		experiment{"ShardedGoMap16", NewShardedGoMap16},
		experiment{"ShardedGoMap32", NewShardedGoMap32},
		experiment{"ShardedGoMap64", NewShardedGoMap64},
	}
)

func BenchmarkControl(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
		}
	})
}

func BenchmarkControl2(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
			rand.Uint32()
		}
	})
}

func BenchmarkControl3(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
		}
	})
}

func BenchmarkControl5(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
		}
	})
}

func BenchmarkControl7(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
			rand.Uint32()
		}
	})
}

func BenchmarkRead(b *testing.B) {
	for _, p := range experiments {
		b.Run(p.label, func(b *testing.B) {
			h := p.newFunc()
			b.StartTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					h.Get(rand.Uint32())
				}
			})
		})
	}
}

func BenchmarkWrite(b *testing.B) {
	for _, p := range experiments {
		b.Run(p.label, func(b *testing.B) {
			h := p.newFunc()
			b.StartTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					h.Put(rand.Uint32(), rand.Uint32())
				}
			})
		})
	}
}

// Equal number of reads and writes.
func BenchmarkReadWrite(b *testing.B) {
	for _, p := range experiments {
		b.Run(p.label, func(b *testing.B) {
			h := p.newFunc()
			b.StartTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					h.Put(rand.Uint32(), rand.Uint32())
					h.Get(rand.Uint32())
				}
			})
		})
	}
}

// Three reads, one write.
func BenchmarkRead3Write1(b *testing.B) {
	for _, p := range experiments {
		b.Run(p.label, func(b *testing.B) {
			h := p.newFunc()
			b.StartTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					h.Put(rand.Uint32(), rand.Uint32())
					h.Get(rand.Uint32())
					h.Get(rand.Uint32())
					h.Get(rand.Uint32())
				}
			})
		})
	}
}

// Three writes, one read.
func BenchmarkRead1Write3(b *testing.B) {
	for _, p := range experiments {
		b.Run(p.label, func(b *testing.B) {
			h := p.newFunc()
			b.StartTimer()
			b.RunParallel(func(pb *testing.PB) {
				for pb.Next() {
					h.Put(rand.Uint32(), rand.Uint32())
					h.Put(rand.Uint32(), rand.Uint32())
					h.Put(rand.Uint32(), rand.Uint32())
					h.Get(rand.Uint32())
				}
			})
		})
	}
}
