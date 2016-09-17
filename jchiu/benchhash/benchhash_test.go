package benchhash

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

type hashPair struct {
	label   string
	newFunc func() HashMap
}

var (
	benchn    = flag.Int("benchn", 100000, "Number of elements to get/put to hash per rep.")
	benchq    = flag.Int("benchq", 10, "Number of goroutines per rep.")
	hashPairs = []hashPair{
		hashPair{"GoMap", NewGoMap},
		hashPair{"GotomicMap", NewGotomicMap},
		hashPair{"ShardedGoMap4", NewShardedGoMap4},
		hashPair{"ShardedGoMap8", NewShardedGoMap8},
		hashPair{"ShardedGoMap16", NewShardedGoMap16},
		hashPair{"ShardedGoMap32", NewShardedGoMap32},
	}
)

func TestMain(m *testing.M) {
	flag.Parse()
	fmt.Printf("n=%d q=%d\n", *benchn, *benchq)
	os.Exit(m.Run())
}

func BenchmarkRead(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiRead(*benchn, *benchq, p.newFunc, b)
		})
	}
}

func BenchmarkWrite(b *testing.B) {
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			MultiWrite(*benchn, *benchq, p.newFunc, b)
		})
	}
}

func benchmarkReadWrite(b *testing.B, fracRead float64) {
	numReadGoRoutines := int(fracRead * float64(*benchq))
	for _, p := range hashPairs {
		b.Run(p.label, func(b *testing.B) {
			ReadWrite(*benchn, numReadGoRoutines, *benchq-numReadGoRoutines,
				p.newFunc, b)
		})
	}
}

func BenchmarkReadWrite1(b *testing.B) { benchmarkReadWrite(b, 0.1) }
func BenchmarkReadWrite3(b *testing.B) { benchmarkReadWrite(b, 0.3) }
func BenchmarkReadWrite5(b *testing.B) { benchmarkReadWrite(b, 0.5) }
func BenchmarkReadWrite7(b *testing.B) { benchmarkReadWrite(b, 0.7) }
func BenchmarkReadWrite9(b *testing.B) { benchmarkReadWrite(b, 0.9) }
