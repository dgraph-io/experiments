package main

import (
	"flag"
	"fmt"
)

var fval = flag.Int("val", 100, "Val")

func findSmallerOrEqualsRecr(ar []int, maxv, left, right int) int {
	if left > right {
		return -1
	}

	pos := (left + right) / 2
	val := ar[pos]
	if val > maxv {
		return findSmallerOrEqualsRecr(ar, maxv, left, pos-1)
	}

	if val == maxv {
		return pos
	}

	tidx := findSmallerOrEqualsRecr(ar, maxv, pos+1, right)
	if tidx == -1 {
		return pos
	}
	return tidx
}

func findSmallerOrEqualsIter(ar []int, maxv int) int {
	left, right := 0, len(ar)-1
	sofar := -1
	for left <= right {
		pos := (left + right) / 2
		val := ar[pos]
		if val > maxv {
			right = pos - 1
			continue
		}

		if val == maxv {
			return pos
		}

		sofar = pos
		left = pos + 1
	}
	return sofar
}

func findSmallerOrEqualsLinear(ar []int, maxv int) int {
	found := -1
	for i := 0; i < len(ar); i++ {
		if ar[i] <= maxv {
			found = i
		} else {
			break
		}
	}
	return found
}

func main() {
	flag.Parse()
	ar := []int{2, 3, 5}
	{
		idx := findSmallerOrEqualsRecr(ar, *fval, 0, len(ar)-1)
		if idx >= 0 {
			fmt.Printf("Idx: %v. Value: %v\n", idx, ar[idx])
		} else {
			fmt.Println("On the left bound")
		}
	}

	{
		i := findSmallerOrEqualsIter(ar, *fval)
		if i >= 0 {
			fmt.Printf("Idx: %v. Value: %v\n", i, ar[i])
		} else {
			fmt.Println("On the left bound")
		}
	}
}
