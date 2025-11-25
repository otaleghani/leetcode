package main

import (
	"cmp"
	"slices"
)

func findMinArrowShots(points [][]int) int {
	slices.SortFunc(points, func(a, b []int) int {
		if n := cmp.Compare(a[1], b[1]); n != 0 {
			return n
		}
		return cmp.Compare(a[0], b[0])
	})

	arrows := 1
	arrow := points[0][1]
	for i := 0; i < len(points); i++ {
		if arrow >= points[i][0] && arrow <= points[i][1] {
			continue
		} else {
			arrow = points[i][1]
			arrows++
		}
	}

	return arrows
}
