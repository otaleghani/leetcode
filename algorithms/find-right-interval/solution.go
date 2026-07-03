package main

import "sort"

type startIdx struct {
	start int
	idx   int
}

func findRightInterval(intervals [][]int) []int {
	n := len(intervals)
	starts := make([]startIdx, n)
	for i, interval := range intervals {
		starts[i] = startIdx{start: interval[0], idx: i}
	}

	sort.Slice(starts, func(i, j int) bool {
		return starts[i].start < starts[j].start
	})

	res := make([]int, n)
	for i, interval := range intervals {
		target := interval[1]
		j := sort.Search(n, func(k int) bool {
			return starts[k].start >= target
		})

		if j < n {
			res[i] = starts[j].idx
		} else {
			res[i] = -1
		}
	}

	return res
}
