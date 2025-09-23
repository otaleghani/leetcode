package main

func findMaxAverage(nums []int, k int) float64 {
	left := 0
	result := 0.0
	currentSum := 0

	for right := 0; right < len(nums); right++ {
		currentSum += nums[right]
		if (right - left + 1) == k {
			currentResult := float64(currentSum) / (float64(right) - float64(left) + 1)
			if currentResult > result || result == 0.0 {
				result = currentResult
			}
			currentSum -= nums[left]
			left++
		}
	}
	return result
}
