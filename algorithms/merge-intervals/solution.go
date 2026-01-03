package main

import "sort"

func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	result := [][]int{}

	result = append(result, intervals[0])

	for i := 1; i < len(intervals); i++ {
		currentStart := intervals[i][0]
		currentEnd := intervals[i][1]

		lastMergedIndex := len(result) - 1
		lastMergedEnd := result[lastMergedIndex][1]

		if currentStart <= lastMergedEnd {
			if currentEnd > lastMergedEnd {
				result[lastMergedIndex][1] = currentEnd
			}
		} else {
			result = append(result, intervals[i])
		}
	}

	return result
}
