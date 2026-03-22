package main

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	max := 1

	for i := range dp {
		dp[i] = 1
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dp[j]+1 > dp[i] {
					dp[i] = dp[j] + 1
				}
			}
		}
		if dp[i] > max {
			max = dp[i]
		}
	}

	return max
}
