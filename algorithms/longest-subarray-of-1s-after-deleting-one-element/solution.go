package main

func longestSubarray(nums []int) int {
	left := 0
	maxLength := 0
	deleted := 1

	for right := range nums {
		if nums[right] == 0 {
			deleted--
		}

		if deleted < 0 {
			if nums[left] == 0 {
				deleted++
			}
			left++
		}

		currentLength := right - left
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
