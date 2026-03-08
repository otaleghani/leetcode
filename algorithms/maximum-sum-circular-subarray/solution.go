package main

func maxSubarraySumCircular(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	totalSum := 0

	maxSum := nums[0]
	curMax := 0

	minSum := nums[0]
	curMin := 0

	for _, num := range nums {
		curMax = max(curMax+num, num)
		maxSum = max(maxSum, curMax)
		curMin = min(curMin+num, num)
		minSum = min(minSum, curMin)
		totalSum += num
	}

	if maxSum < 0 {
		return maxSum
	}

	return max(maxSum, totalSum-minSum)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
