package main

import (
	"crypto/rand"
	"io/ioutil"
	"testing"
)

func benchGobEncode(b *testing.B, sz int) {
	buf := make([]byte, sz)
	_, err := rand.Read(buf)
	if err != nil {
		b.Error(err)
		b.Fail()
	}
	q := new(Query)
	q.Data = buf
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := GobEncode(q, ioutil.Discard); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkGobEncode_1K(b *testing.B)  { benchGobEncode(b, 1000) }
func BenchmarkGobEncode_1M(b *testing.B)  { benchGobEncode(b, 1000000) }
func BenchmarkGobEncode_50M(b *testing.B) { benchGobEncode(b, 50000000) }

func benchEncode(b *testing.B, sz int) {
	buf := make([]byte, sz)
	_, err := rand.Read(buf)
	if err != nil {
		b.Error(err)
		b.Fail()
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		if err := Encode(buf, ioutil.Discard); err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkEncode_1K(b *testing.B)  { benchEncode(b, 1000) }
func BenchmarkEncode_1M(b *testing.B)  { benchEncode(b, 1000000) }
func BenchmarkEncode_50M(b *testing.B) { benchEncode(b, 50000000) }
