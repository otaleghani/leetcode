package main

import "math"

func minSubArrayLen(target int, nums []int) int {
	left := 0
	sum := 0
	min := math.MaxInt

	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for target <= sum {
			currLen := right - left + 1
			if currLen < min {
				min = currLen
			}
			sum -= nums[left]
			left++
		}
	}

	if min == math.MaxInt {
		return 0
	}
	return min
}
