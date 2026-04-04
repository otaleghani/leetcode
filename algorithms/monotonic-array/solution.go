package main

func isMonotonic(nums []int) bool {
	if len(nums) < 1 {
		return true
	}

	last := nums[0]
	isAsc := true
	j := 1
	for j < len(nums) && last == nums[j] {
		j++
	}
	if j == len(nums) {
		return true
	}
	if last < nums[j] {
		isAsc = false
	}
	for i := 1; i < len(nums); i++ {
		if isAsc {
			if last < nums[i] {
				return false
			}
			last = nums[i]
		} else {
			if last > nums[i] {
				return false
			}
			last = nums[i]
		}
	}

	return true
}
