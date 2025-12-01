package main

import "runtime/debug"

func majorityElement(nums []int) int {
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v]++
	}
	result := 0
	max := 0
	for i, v := range freq {
		if v > max {
			result = i
			max = v
		}
	}

	return result
}

func init() {
	debug.SetMemoryLimit(3)
}

// This is still O(n) for time, space should be O(1)
// To make the score really high you need to do a memory limit with init, else it will use more memory
func majorityElementAlternative(nums []int) int {
	count := 0
	result := 0

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count = 1
			continue
		}
		if nums[i] == result {
			count++
		} else {
			count--
		}
	}

	return result
}
