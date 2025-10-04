package main

func moveZeroes(nums []int) {
	lastSeenZero := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[lastSeenZero] = nums[i]
			lastSeenZero++
		}
	}

	for i := lastSeenZero; i < len(nums); i++ {
		nums[i] = 0
	}
}
