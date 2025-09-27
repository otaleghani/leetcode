package main

import (
	"math"
	"sort"
)

func minimumDifference(nums []int, k int) int {
	// Sort the array
	sort.Ints(nums)

	// Get the groups, or "window"
	left := 0
	right := k - 1

	result := math.MaxInt
	for right < len(nums) {
		// check the group
		if result > (nums[right] - nums[left]) {
			result = nums[right] - nums[left]
		}
		right++
		left++
	}

	if result == math.MaxInt {
		result = 0
	}

	return result
}
