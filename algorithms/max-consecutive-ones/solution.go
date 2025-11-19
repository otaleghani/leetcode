package main

func findMaxConsecutiveOnes(nums []int) int {
	max := 0
	currMax := 0
	for _, v := range nums {
		if v == 1 {
			currMax += 1
			if currMax > max {
				max = currMax
			}
		} else {
			currMax = 0
		}
	}

	return max
}
