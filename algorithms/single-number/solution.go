package main

func singleNumber(nums []int) int {
	var result = 0

	for i := range len(nums) {
		result = result ^ nums[i]
	}

	return result
}

func singleNumberAlt(nums []int) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}
