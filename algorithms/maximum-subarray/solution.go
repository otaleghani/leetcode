package main

func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	curr := nums[0]
	maxSum := nums[0]

	for i := 1; i < len(nums); i++ {
		curr = max(nums[i], curr+nums[i])
		maxSum = max(maxSum, curr)
	}

	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
