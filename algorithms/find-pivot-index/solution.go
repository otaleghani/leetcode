package main

func pivotIndex(nums []int) int {
	prefix := make([]int, len(nums))
	suffix := make([]int, len(nums))

	prefix[0] = 0
	suffix[len(suffix)-1] = 0

	for i := 1; i < len(nums); i++ {
		prefix[i] = prefix[i-1] + nums[i-1]
		suffix[len(suffix)-1-i] = suffix[len(suffix)-i] + nums[len(nums)-i]
	}

	for i := 0; i < len(prefix); i++ {
		if prefix[i] == suffix[i] {
			return i
		}
	}

	return -1
}
