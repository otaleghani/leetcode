package main

func minimumOperations(nums []int) int {
	ops := 0
	for _, val := range nums {
		ops += min(val%3, 3-(val%3))
	}
	return ops
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
