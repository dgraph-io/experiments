package benchhash2

import (
	"math/rand"
	"strconv"
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

func BenchmarkReadControl(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
		}
	})
}

func BenchmarkWriteControl(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			rand.Uint32()
			rand.Uint32()
		}
	})
}

func benchmarkReadWriteControl(b *testing.B, readFrac float32) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			if rand.Float32() < readFrac {
				rand.Uint32()
			} else {
				rand.Uint32()
				rand.Uint32()
			}
		}
	})
}

func BenchmarkReadWriteControl(b *testing.B) {
	for i := 0; i <= 20; i++ {
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			benchmarkReadWriteControl(b, float32(i)/20)
		})
	}
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

// readFrac is fraction of operations that are reads.
func BenchmarkReadWrite(b *testing.B) {
	for i := 0; i <= 20; i++ {
		readFrac := float32(i) / 20.0
		b.Run(strconv.Itoa(i), func(b *testing.B) {
			for _, p := range experiments {
				b.Run(p.label, func(b *testing.B) {
					h := p.newFunc()
					b.StartTimer()
					b.RunParallel(func(pb *testing.PB) {
						for pb.Next() {
							if rand.Float32() < readFrac {
								h.Get(rand.Uint32())
							} else {
								h.Put(rand.Uint32(), rand.Uint32())
							}
						}
					})
				})
			}
		})
	}
}
