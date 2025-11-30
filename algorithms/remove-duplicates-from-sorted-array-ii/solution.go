package main

func removeDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}

	left := 2
	for right := 2; right < len(nums); right++ {
		if nums[right] != nums[left-1] || nums[right] != nums[left-2] {
			nums[left] = nums[right]
			left++
		}
	}

	return left
}
