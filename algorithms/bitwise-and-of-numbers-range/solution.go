package main

func rangeBitwiseAnd(left int, right int) int {
	shiftCount := 0
	for left < right {
		left >>= 1
		right >>= 1
		shiftCount++
	}
	return left << shiftCount
}
