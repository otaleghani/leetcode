package main

func minOperations(nums []int) int {
	stack := []int{}
	result := 0
	for _, val := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > val {
			stack = stack[:len(stack)-1] // Pop the num from the stack
		}
		if val == 0 {
			continue
		}
		if len(stack) == 0 || stack[len(stack)-1] < val {
			result++
			stack = append(stack, val)
		}
	}
	return result
}
