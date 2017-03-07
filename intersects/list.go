package intersect

import (
	"fmt"
	"sort"
)

func mergeIntersect(a, b []uint64, final *[]uint64) {
	ma, mb := len(a), len(b)
	i, j := 0, 0
	for i < ma && j < mb {
		if a[i] == b[j] {
			*final = append(*final, a[i])
			i++
			j++

		} else if a[i] < b[j] {
			for i = i + 1; i < ma && a[i] < b[j]; i++ {
			}

		} else {
			for j = j + 1; j < mb && a[i] > b[j]; j++ {
			}
		}

	}
}

func mergeDeltaIntersect(a, b *DeltaList, final *[]uint64) {
	if len(a.Uids) > len(b.Uids) {
		panic("this is wrong")
	}
	if len(a.Uids) == 0 || len(b.Uids) == 0 {
		return
	}

	var ea uint64
	bsize := int(b.BucketSize)
	var bidx int
	var eb uint64 = b.Uids[0]

	bucketIdx := 0
	numb := len(b.Buckets)
	numbu := len(b.Uids)
	end := bsize
	endBucket := b.Buckets[0]

OUTER:
	for _, da := range a.Uids {
		ea += da

		// find the bucket
		for ; endBucket < ea && bucketIdx < numb; bucketIdx++ {
			bidx = bsize * bucketIdx
			endBucket = b.Buckets[bucketIdx]

			if bucketIdx > 0 {
				eb = b.Buckets[bucketIdx-1] + b.Uids[bidx]
			} else {
				eb = b.Uids[bidx]
			}

			end = bsize * (bucketIdx + 1)
			if end > numbu {
				end = numbu
			}
			fmt.Printf("eb=%v endbucket=%v\n", eb, endBucket)
		}
		if ea > endBucket {
			break
		}

		// Got the right bucket.
		for {
			fmt.Printf("ea=%d eb=%d endbucket=%d\n", ea, eb, endBucket)
			if eb >= ea {
				if eb == ea {
					*final = append(*final, ea)
					fmt.Printf("OUT=%v\n", ea)
				}
				break
			}
			bidx++
			if bidx >= end {
				break
			}
			eb += b.Uids[bidx]
		}
		if bidx == numbu {
			break OUTER
		}
	}
}

func BinIntersect(d, q []uint64, final *[]uint64) {
	ld := len(d)
	lq := len(q)
	if ld == 0 || lq == 0 || d[ld-1] < q[0] || q[lq-1] < d[0] {
		return
	}
	if ld < lq {
		panic("what")
	}

	val := d[0]
	minq := sort.Search(len(q), func(i int) bool {
		return q[i] >= val
	})

	val = d[len(d)-1]
	maxq := sort.Search(len(q), func(i int) bool {
		return q[i] >= val
	})

	binIntersect(d, q[minq:maxq], final)
}

func binIntersect(d, q []uint64, final *[]uint64) {
	if len(d) == 0 || len(q) == 0 {
		return
	}
	midq := len(q) / 2
	qval := q[midq]
	midd := sort.Search(len(d), func(i int) bool {
		return d[i] >= qval
	})

	dd := d[0:midd]
	qq := q[0:midq]
	if len(dd) > len(qq) { // D > Q
		binIntersect(dd, qq, final)
	} else {
		binIntersect(qq, dd, final)
	}

	if midd >= len(d) {
		return
	}
	if d[midd] == qval {
		*final = append(*final, qval)
	} else {
		midd -= 1
	}

	dd = d[midd+1:]
	qq = q[midq+1:]
	if len(dd) > len(qq) { // D > Q
		binIntersect(dd, qq, final)
	} else {
		binIntersect(qq, dd, final)
	}
}

type S struct {
	d []uint64
	q []uint64
}

func binIterative(d, q []uint64, final *[]uint64) {
	stack := make([]S, 0, 100)
	stack = append(stack, S{d, q})

	for len(stack) > 0 {
		l := len(stack) - 1
		s := stack[l]
		stack = stack[:l] // pop

		if len(s.d) > len(s.q) {
			d = s.d
			q = s.q
		} else {
			d = s.q
			q = s.d
		}

		if len(d) == 0 || len(q) == 0 {
			continue
		}

		midq := len(q) / 2
		qval := q[midq]
		midd := sort.Search(len(d), func(i int) bool {
			return d[i] >= qval
		})
		stack = append(stack, S{d[0:midd], q[0:midq]})

		if midd >= len(d) {
			continue
		}
		if d[midd] == qval {
			*final = append(*final, qval)
		} else {
			midd -= 1
		}
		stack = append(stack, S{d[midd+1:], q[midq+1:]})
	}
	//sort.Slice(*final, func(i, j int) bool {
	//return (*final)[i] < (*final)[j]
	//})
}

func encodeDelta(d []uint64, bucketSize int) *DeltaList {
	l := new(DeltaList)
	if len(d) == 0 {
		return l
	}
	l.BucketSize = int32(bucketSize)

	var last uint64
	for i, cur := range d {
		if i%bucketSize == bucketSize-1 { // Store the max of the bucket.
			l.Buckets = append(l.Buckets, cur)
		}
		l.Uids = append(l.Uids, cur-last)
		last = cur
	}
	if len(l.Buckets) == 0 || l.Buckets[len(l.Buckets)-1] != last {
		l.Buckets = append(l.Buckets, last)
	}

	return l
}

func encodeFixed(d []uint64) *FixedList {
	f := new(FixedList)
	f.Uids = d
	return f
}
