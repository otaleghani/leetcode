package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}

	left, right := 0, len(nums)-1
	for left <= right {
		if left == right && nums[left] == val {
			right--
			left++
			continue
		}
		if nums[left] == val {
			// Is right a swappable value?
			for nums[right] == val && right > 0 {
				right--
			}
			if left >= right {
				continue
			}
			nums[left], nums[right] = nums[right], nums[left]
			for nums[right] == val && right > 0 {
				right--
			}
		}
		left++
	}

	return right + 1
}
