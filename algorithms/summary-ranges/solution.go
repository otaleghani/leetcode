package main

import "fmt"

func summaryRanges(nums []int) []string {
	result := []string{}
	if len(nums) == 0 {
		return result
	}
	start := 0
	count := nums[start]

	for i := 0; i < len(nums); i++ {
		if nums[i] == count {
			count++
			continue
		}
		if start == i-1 {
			result = append(result, fmt.Sprintf("%d", nums[start]))
			start = i
			count = nums[i] + 1
			continue
		}
		result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
		start = i
		count = nums[i] + 1
	}

	if start == len(nums)-1 {
		result = append(result, fmt.Sprintf("%d", nums[start]))
	} else {
		result = append(result, fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]))
	}

	return result
}
