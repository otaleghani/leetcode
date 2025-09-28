package main

import "strconv"

func divisorSubstrings(num int, k int) int {
	result := 0
	digits := []int{}
	oldNum := num

	if num == 0 {
		digits = append(digits, 0)
	} else {
		for num > 0 {
			current := num % 10
			digits = append([]int{current}, digits...)
			num /= 10
		}
	}

	left := 0
	for right := k; right <= len(digits); right++ {
		combinedNumber := 0
		for _, digit := range digits[left:right] {
			combinedNumber *= 10
			combinedNumber += digit
		}
		if combinedNumber != 0 && oldNum%combinedNumber == 0 {
			result++
		}

		left++
	}
	return result
}

func divisorSubstringsAlternative(num int, k int) int {
	result := 0
	s := strconv.Itoa(num)

	for i := 0; i <= len(s)-k; i++ {
		sub := s[i : i+k]
		val, _ := strconv.Atoi(sub)
		if val != 0 && num%val == 0 {
			result++
		}
	}
	return result
}
