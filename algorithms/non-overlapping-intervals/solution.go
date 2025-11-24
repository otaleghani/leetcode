package main

import (
	"cmp"
	"math"
	"slices"
)

func eraseOverlapIntervals(intervals [][]int) int {
	// Sort the intervals by their last point and then their first point
	slices.SortFunc(intervals, func(a, b []int) int {
		if n := cmp.Compare(a[1], b[1]); n != 0 {
			return n
		}

		return cmp.Compare(a[0], b[0])
	})

	dp := make([]int, len(intervals))
	for i := range dp {
		dp[i] = math.MinInt32
	}

	var f func(currentSkipped, pos int, currList [][]int) int
	f = func(currentSkipped, pos int, currList [][]int) int {
		if currentSkipped >= len(intervals) || pos >= len(intervals) {
			return currentSkipped
		}

		if dp[pos] != math.MinInt32 {
			return dp[pos]
		}

		resAdded := math.MaxInt32
		resSkipped := math.MaxInt32
		if len(currList) == 0 {
			resAdded = f(currentSkipped, pos+1, append(currList, intervals[pos]))
			resSkipped = f(currentSkipped+1, pos+1, currList)
			dp[pos] = min(resAdded, resSkipped)
			return dp[pos]
		}

		// Check if valid
		last := currList[len(currList)-1]
		if !(last[1] <= intervals[pos][1] && last[1] > intervals[pos][0]) {
			resAdded = f(currentSkipped, pos+1, append(currList, intervals[pos]))
		}
		resSkipped = f(currentSkipped+1, pos+1, currList)

		dp[pos] = min(resAdded, resSkipped)
		return dp[pos]
	}

	list := make([][]int, 0)
	return f(0, 0, list)
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// Alternative version passing just the last element to the dynamic programming function and not the hole array
func eraseOverlapIntervalsAlt(intervals [][]int) int {
	// Sort the intervals by their last point and then their first point
	slices.SortFunc(intervals, func(a, b []int) int {
		if n := cmp.Compare(a[1], b[1]); n != 0 {
			return n
		}

		return cmp.Compare(a[0], b[0])
	})

	dp := make([]int, len(intervals))
	for i := range dp {
		dp[i] = math.MinInt32
	}

	var f func(currentSkipped, pos int, last []int) int
	f = func(currentSkipped, pos int, last []int) int {
		if currentSkipped >= len(intervals) || pos >= len(intervals) {
			return currentSkipped
		}

		if dp[pos] != math.MinInt32 {
			return dp[pos]
		}

		resAdded := math.MaxInt32
		resSkipped := math.MaxInt32
		if len(last) == 0 {
			resAdded = f(currentSkipped, pos+1, intervals[pos])
			resSkipped = f(currentSkipped+1, pos+1, last)
			dp[pos] = min(resAdded, resSkipped)
			return dp[pos]
		}

		// Check if valid
		if !(last[1] <= intervals[pos][1] && last[1] > intervals[pos][0]) {
			resAdded = f(currentSkipped, pos+1, intervals[pos])
		}
		resSkipped = f(currentSkipped+1, pos+1, last)

		dp[pos] = min(resAdded, resSkipped)
		return dp[pos]
	}

	list := []int{}
	return f(0, 0, list)
}

// The better approach (greedy)
func eraseOverlapIntervalsGreedy(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	slices.SortFunc(intervals, func(a, b []int) int {
		return cmp.Compare(a[1], b[1])
	})

	end := intervals[0][1]
	count := 0

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] < end {
			count++
		} else {
			end = intervals[i][1]
		}
	}

	return count
}
