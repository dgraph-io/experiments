package main

import (
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"testing"

	"github.com/boltdb/bolt"
)

func benchWriteBolt(b *testing.B, N int) {
	dir, err := ioutil.TempDir("", "bolt")
	if err != nil {
		b.Error(err)
		return
	}
	defer os.RemoveAll(dir)

	db, err := bolt.Open(dir+"/bolt.db", 0600, nil)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := []byte(strconv.Itoa(i))
		if err := writeNBytes(db, k, N); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkWriteBolt_1024(b *testing.B)  { benchWriteBolt(b, 1024) }
func BenchmarkWriteBolt_10KB(b *testing.B)  { benchWriteBolt(b, 10240) }
func BenchmarkWriteBolt_500KB(b *testing.B) { benchWriteBolt(b, 1<<19) }
func BenchmarkWriteBolt_1MB(b *testing.B)   { benchWriteBolt(b, 1<<20) }

func benchReadBolt(b *testing.B, N int) {
	dir, err := ioutil.TempDir("", "bolt")
	if err != nil {
		b.Error(err)
		return
	}
	defer os.RemoveAll(dir)

	numKeys := 100
	db, err := bolt.Open(dir+"/bolt.db", 0600, nil)
	for i := 0; i < numKeys; i++ {
		k := []byte(strconv.Itoa(i))
		if err := writeNBytes(db, k, N); err != nil {
			b.Error(err)
		}
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		k := rand.Int() % numKeys
		key := []byte(strconv.Itoa(k))
		n := readValue(db, key)
		if n != N {
			b.Errorf("Expected: %v.  Got: %v", N, n)
		}
	}
}

func BenchmarkReadBolt_1024(b *testing.B)  { benchReadBolt(b, 1024) }
func BenchmarkReadBolt_10KB(b *testing.B)  { benchReadBolt(b, 10240) }
func BenchmarkReadBolt_500KB(b *testing.B) { benchReadBolt(b, 1<<19) }
func BenchmarkReadBolt_1MB(b *testing.B)   { benchReadBolt(b, 1<<20) }
