package main

import "math"

func maxSubarraySum(nums []int, k int) int64 {
	minPrefix := make([]int64, k)
	for i := range minPrefix {
		minPrefix[i] = math.MaxInt64
	}

	minPrefix[0] = 0
	var currSum int64 = 0
	var maxResult int64 = math.MinInt64
	foundValid := false

	for i, v := range nums {
		currSum += int64(v)
		currentRem := (i + 1) % k
		if minPrefix[currentRem] != math.MaxInt64 {
			candidate := currSum - minPrefix[currentRem]
			if candidate > maxResult {
				maxResult = candidate
			}
			foundValid = true
		}
		if currSum < minPrefix[currentRem] {
			minPrefix[currentRem] = currSum
		}
	}

	if !foundValid {
		return 0
	}

	return maxResult
}
