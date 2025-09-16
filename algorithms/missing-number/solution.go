package main

func missingNumber(nums []int) int {
	result := 0

	for i := range nums {
		result ^= nums[i]
	}
	for i := range len(nums) + 1 {
		result ^= i
	}

	return result
}
