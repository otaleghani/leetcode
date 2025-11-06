package main

func rob(nums []int) int {
	rob1, rob2 := 0, 0

	for _, n := range nums {
		temp := max(n+rob2, rob1)
		rob2 = rob1
		rob1 = temp
	}

	return rob1
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
