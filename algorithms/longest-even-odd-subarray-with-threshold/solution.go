package main

func longestAlternatingSubarray(nums []int, threshold int) int {
	// left has to be even
	// Numbers have to be even-odd-even
	// nums[i] <= threshold

	result := 0
	for left := 0; left < len(nums); left++ {
		if nums[left]%2 == 0 && nums[left] <= threshold {
			if left == len(nums)-1 {
				// we are at the end of the array
				currentResult := 1
				if currentResult > result {
					result = currentResult
				}
			}

			// We can start testing
			for right := left + 1; right < len(nums); right++ {
				if nums[right] > threshold {
					// count till here
					currentResult := right - left
					if currentResult > result {
						result = currentResult
					}
					break
				}
				if nums[right]%2 == nums[right-1]%2 {
					// The current one is not alternating, so stop one before
					currentResult := right - left
					if currentResult > result {
						result = currentResult
					}
					break
				}
				if right == len(nums)-1 {
					// we are at the end of the array
					currentResult := right + 1 - left
					if currentResult > result {
						result = currentResult
					}
					break
				}
			}
		}
	}
	return result
}
