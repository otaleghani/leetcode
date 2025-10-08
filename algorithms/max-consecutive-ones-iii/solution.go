package main

func longestOnes(nums []int, k int) int {
	left := 0
	right := 0
	max := -1
	currentK := 0
	for right < len(nums) {
		if nums[right] == 1 {
			max = calcMax(max, right-left)
			right++
		} else {
			if k == 0 {
				right++
				left = right
				continue
			}
			currentK++
			if currentK > k {
				// overflow, move window until we find another 0!
				for nums[left] == 1 {
					left++
				}
				currentK--
				left++
			}

			max = calcMax(max, right-left)
			right++
		}
	}

	return max + 1
}

func calcMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// Better solution with a clear window
func longestOnesAlternative(nums []int, k int) int {
	left := 0
	maxLength := 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			k--
		}

		if k < 0 {
			if nums[left] == 0 {
				k++
			}
			left++
		}

		currentLength := right - left + 1
		if currentLength > maxLength {
			maxLength = currentLength
		}
	}

	return maxLength
}
