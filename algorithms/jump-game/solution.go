package main

func canJump(nums []int) bool {
	i := 0
	maxExplore := 0
	lastElement := len(nums) - 1
	for i <= maxExplore {
		currMaxRange := i + nums[i] // Current position + the max jump in this position
		if currMaxRange > maxExplore {
			maxExplore = currMaxRange
		}
		if maxExplore >= lastElement {
			return true
		}
		i++
	}
	return false
}
