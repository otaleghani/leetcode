package main

func productExceptSelf(nums []int) []int {
	prefix := make(map[int]int)
	suffix := make(map[int]int)

	prefix[-1] = 1
	suffix[len(nums)] = 1

	for i := range len(nums) {
		prefix[i] = prefix[i-1] * nums[i]
		suffix[len(nums)-i-1] = suffix[len(nums)-i] * nums[len(nums)-i-1]
	}

	var result []int

	for i := range len(nums) {
		result = append(result, prefix[i-1]*suffix[i+1])
	}

	return result
}
