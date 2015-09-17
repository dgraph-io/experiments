package convert

import (
	"bytes"
	"hash"
	"hash/crc64"
	"io"
	"math/rand"
	"testing"
	"time"

	"github.com/spaolacci/murmur3"
)

func TestUseCrc(t *testing.T) {
	table := crc64.MakeTable(crc64.ISO)
	h := crc64.New(table)

	input := "someUniqueId"
	io.WriteString(h, input)
	sum1 := h.Sum64()

	h.Reset()
	input = "someOtherId"
	io.WriteString(h, input)
	sum2 := h.Sum64()

	if sum1 == sum2 {
		t.Errorf("Sums shouldn't match [%x] [%x]\n", sum1, sum2)
		t.Fail()
		return
	}

	h.Reset()
	input = "someUniqueId"
	io.WriteString(h, input)
	sum3 := h.Sum64()

	if sum1 != sum3 {
		t.Errorf("Sums should match [%x] [%x]\n", sum1, sum3)
		t.Fail()
		return
	}
}

const alphachars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

var r = rand.New(rand.NewSource(time.Now().UnixNano()))

// UniqueString generates a unique string only using the characters from
// alphachars constant, with length as specified.
func uniqueString(alpha int) string {
	var buf bytes.Buffer
	for i := 0; i < alpha; i++ {
		idx := r.Intn(len(alphachars))
		buf.WriteByte(alphachars[idx])
	}
	return buf.String()
}

func getUids(num int) []string {
	l := make([]string, num)
	for i := 0; i < num; i++ {
		l[i] = uniqueString(10)
	}
	return l
}

const uidSize = 10000000 // At 10M, this would take over 100MB of RAM per test.
func testCollissions(t *testing.T, h hash.Hash64) {
	uids := getUids(uidSize)
	results := make(map[uint64]bool)
	cols := 0

	for i := 0; i < uidSize; i++ {
		h.Reset()
		io.WriteString(h, uids[i])
		s := h.Sum64()
		if _, col := results[s]; col {
			cols += 1
		} else {
			results[s] = true
		}
	}
	if cols > 0 {
		t.Errorf("Found %v collissions for uidSize %v\n", cols, uidSize)
	}
}

func TestUseCrc_ISOCollissions(t *testing.T) {
	table := crc64.MakeTable(crc64.ISO)
	h := crc64.New(table)
	testCollissions(t, h)
}

func TestUseCrc_ECMACollissions(t *testing.T) {
	table := crc64.MakeTable(crc64.ECMA)
	h := crc64.New(table)
	testCollissions(t, h)
}

func TestUseMurmur_Collissions(t *testing.T) {
	h := murmur3.New64()
	testCollissions(t, h)
}

var result uint64

func benchmarkHash(b *testing.B, h hash.Hash64) {
	uids := getUids(b.N)
	var s uint64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		h.Reset()
		io.WriteString(h, uids[i])
		s = h.Sum64()
	}
	result = s
}

func BenchmarkUseCrc_ISO(b *testing.B) {
	table := crc64.MakeTable(crc64.ISO)
	h := crc64.New(table)
	benchmarkHash(b, h)
}

func BenchmarkUseCrc_ECMA(b *testing.B) {
	table := crc64.MakeTable(crc64.ECMA)
	h := crc64.New(table)
	benchmarkHash(b, h)
}

func BenchmarkUseMurmur(b *testing.B) {
	h := murmur3.New64()
	benchmarkHash(b, h)
}
