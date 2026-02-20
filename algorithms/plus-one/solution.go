package main

func plusOne(digits []int) []int {
	i := len(digits) - 1
	if digits[i] != 9 {
		digits[i] += 1
		return digits
	}

	for i >= 0 && digits[i] == 9 {
		digits[i] = 0
		i--
	}

	if digits[0] == 0 {
		digits = append([]int{1}, digits...)
	} else {
		digits[i] += 1
	}

	return digits
}
