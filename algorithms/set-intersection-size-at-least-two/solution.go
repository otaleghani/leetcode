package main

import (
	"cmp"
	"slices"
)

func intersectionSizeTwo(intervals [][]int) int {
	slices.SortFunc(intervals, func(a, b []int) int {
		if n := cmp.Compare(a[1], b[1]); n != 0 {
			return n
		}
		return cmp.Compare(b[0], a[0])
	})

	result := 0
	p1 := -1
	p2 := -1

	for _, interval := range intervals {
		left := interval[0]
		right := interval[1]

		if p2 < left {
			result += 2
			p1 = right - 1
			p2 = right
		} else if p1 < left {
			result += 1
			p1 = p2
			p2 = right
		}
	}

	return result
}
