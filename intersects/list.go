package intersect

import (
	"fmt"
	"sort"
)

func merge(a, b []uint64) []uint64 {
	ma, mb := len(a), len(b)
	i, j := 0, 0
	out := make([]uint64, 0, 100)
	for i < ma && j < mb {
		if a[i] == b[j] {
			out = append(out, a[i])
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
	return out
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
		fmt.Printf("Adding: %d\n", qval)
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
