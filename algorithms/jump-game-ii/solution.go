package main

func jump(nums []int) int {
	if len(nums)-1 <= 0 {
		return 0
	}

	minJumps := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i]+i >= len(nums)-1 {
			return minJumps[i] + 1
		}
		for j := range nums[i] {
			if i+j+1 < len(nums) && minJumps[i+j+1] == 0 {
				minJumps[i+j+1] = minJumps[i] + 1
			}
		}
	}

	return 0
}
