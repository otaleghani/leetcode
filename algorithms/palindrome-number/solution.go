package main

import (
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	str = string(runes)

	reversed, _ := strconv.Atoi(str)

	return x == reversed
}

func isPalindromeAlt(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = (revertedNumber * 10) + (x % 10)
		x /= 10
	}

	return x == revertedNumber || x == revertedNumber/10
}
