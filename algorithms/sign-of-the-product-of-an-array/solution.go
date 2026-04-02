package main

func arraySign(nums []int) int {
	res := signFunc(nums[0])
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		}

		res *= signFunc(nums[i])
	}

	return res
}

func signFunc(x int) int {
	if x > 0 {
		return 1
	} else if x == 0 {
		return 0
	}
	return -1
}
