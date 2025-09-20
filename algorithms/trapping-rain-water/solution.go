package main

func trap(height []int) int {
	if len(height) < 3 {
		return 0
	}

	result := 0
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0

	for left < right {
		if height[left] < height[right] {
			if height[left] >= maxLeft {
				maxLeft = height[left]
			} else {
				result += maxLeft - height[left]
			}
			left++
		} else {
			if height[right] >= maxRight {
				maxRight = height[right]
			} else {
				result += maxRight - height[right]
			}
			right--
		}
	}

	return result
}
