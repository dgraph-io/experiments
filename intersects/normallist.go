package intersect

import (
	"sort"

	"github.com/dgraph-io/dgraph/task"
)

func IntersectWith(u, v *task.List) {
	n := len(u.Uids)
	m := len(v.Uids)

	if n > m {
		n, m = m, n
	}
	if n == 0 {
		n += 1
	}
	// Select appropriate function based on heuristics.
	ratio := float64(m) / float64(n)
	if ratio < 100 {
		IntersectWithLin(u, v)
	} else if ratio < 500 {
		IntersectWithJump(u, v)
	} else {
		IntersectWithBin(u, v)
	}
}

// IntersectWith intersects u with v. The update is made to u.
// u, v should be sorted.
func IntersectWithLin(u, v *task.List) {
	out := u.Uids[:0]
	n := len(u.Uids)
	m := len(v.Uids)
	for i, k := 0, 0; i < n && k < m; {
		uid := u.Uids[i]
		vid := v.Uids[k]
		if uid > vid {
			for k = k + 1; k < m && v.Uids[k] < uid; k++ {
			}
		} else if uid == vid {
			out = append(out, uid)
			k++
			i++
		} else {
			for i = i + 1; i < n && u.Uids[i] < vid; i++ {
			}
		}
	}
	//u.Uids = out
}

func IntersectWithJump(u, v *task.List) {
	out := u.Uids[:0]
	n := len(u.Uids)
	m := len(v.Uids)
	jump := 30
	for i, k := 0, 0; i < n && k < m; {
		uid := u.Uids[i]
		vid := v.Uids[k]
		if uid == vid {
			out = append(out, uid)
			k++
			i++
		} else if k+jump < m && uid > v.Uids[k+jump] {
			k = k + jump
		} else if i+jump < n && vid > u.Uids[i+jump] {
			i = i + jump
		} else if uid > vid {
			for k = k + 1; k < m && v.Uids[k] < uid; k++ {
			}
		} else {
			for i = i + 1; i < n && u.Uids[i] < vid; i++ {
			}
		}
	}
	//u.Uids = out
}

func IntersectWithBin(u, v *task.List) {
	out := u.Uids[:0]
	m := len(u.Uids)
	n := len(v.Uids)
	// We want to do binary search on bigger list.
	smallList, bigList := u.Uids, v.Uids
	if m > n {
		smallList, bigList = bigList, smallList
	}
	// This is reduce the search space after every match.
	searchList := bigList
	for _, uid := range smallList {
		idx := sort.Search(len(searchList), func(i int) bool {
			return searchList[i] >= uid
		})
		if idx < len(searchList) && searchList[idx] == uid {
			out = append(out, uid)
			// The next UID would never be at less than this idx
			// as the list is sorted.
			searchList = searchList[idx:]
		}
	}
	//u.Uids = out
}
