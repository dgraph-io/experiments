package x

import (
	"bytes"
	"fmt"
	"math/rand"
	"sort"
)

const alphachars = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func UniqueString(alpha int) string {
	var buf bytes.Buffer
	for i := 0; i < alpha; i++ {
		idx := rand.Intn(len(alphachars))
		buf.WriteByte(alphachars[idx])
	}
	return buf.String()
}

type int64arr []int64

func (a int64arr) Len() int           { return len(a) }
func (a int64arr) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a int64arr) Less(i, j int) bool { return a[i] < a[j] }

func SortedInt(alpha int) []int64 {
	result := make([]int64, alpha)
	for i := 0; i < alpha; i++ {
		result[i] = rand.Int63()
	}
	sort.Sort(int64arr(result))
	return result
}

func SortedString(alpha int) []string {
	result := make([]string, alpha)
	for i := 0; i < alpha; i++ {
		result[i] = UniqueString(rand.Intn(11))
	}
	sort.Sort(sort.StringSlice(result))
	return result
}

func linearSearchInt(arr []int64, pos int, n int64) int {
	for i := pos; i < len(arr); i++ {
		if arr[i] > n {
			return i - 1
		}
	}
	return len(arr) - 1
}

func findIndexInt(arr []int64, n int64) int {
	i := 0
	j := len(arr) - 1

	for n > arr[i] {
		if n > arr[j] {
			// The right limit found is already lower than n.
			// Let's do linear search here.
			return linearSearchInt(arr, j, n)
		}

		mid := (i + j) / 2
		if mid == i {
			return linearSearchInt(arr, i, n)
		} else if n > arr[mid] {
			i = mid + 1
		} else if n < arr[mid] {
			j = mid - 1
		} else {
			return linearSearchInt(arr, mid, n)
		}
	}
	return i
}

func mergeInt(arr []int64, n int64) []int64 {
	arr = append(arr, n)
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		} else {
			break
		}
	}
	return arr
}

func linearSearchString(arr []string, pos int, n string) int {
	for i := pos; i < len(arr); i++ {
		if arr[i] > n {
			return i - 1
		}
	}
	return len(arr) - 1
}

func findIndexString(arr []string, n string) int {
	i := 0
	j := len(arr) - 1
	for n > arr[i] {
		if arr[j] < n {
			return linearSearchString(arr, j, n)
		}

		mid := (i + j) / 2
		if n > arr[mid] {
			i = mid + 1
		} else if n < arr[mid] {
			j = mid - 1
		} else {
			//  n == arr[mid]
			// linear search
			return linearSearchString(arr, mid, n)
		}
	}
	return i
}

func PrintList(l []int64) {
	for i := 0; i < len(l); i++ {
		fmt.Printf("pos=[%v] val=[%v]\n", i, l[i])
	}
	fmt.Println("=============")
}

func mergeString(arr []string, n string) []string {
	arr = append(arr, n)
	for i := len(arr) - 1; i > 0; i-- {
		if arr[i] < arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		} else {
			break
		}
	}
	return arr
}
