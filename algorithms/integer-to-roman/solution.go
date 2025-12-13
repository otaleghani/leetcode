package main

import "strings"

func intToRoman(num int) string {
	roman := make(map[int]string)
	keys := make([]int, 13)

	roman[0] = "I"
	roman[1] = "IV"
	roman[2] = "V"
	roman[3] = "IX"
	roman[4] = "X"
	roman[5] = "XL"
	roman[6] = "L"
	roman[7] = "XC"
	roman[8] = "C"
	roman[9] = "CD"
	roman[10] = "D"
	roman[11] = "CM"
	roman[12] = "M"

	keys[0] = 1
	keys[1] = 4
	keys[2] = 5
	keys[3] = 9
	keys[4] = 10
	keys[5] = 40
	keys[6] = 50
	keys[7] = 90
	keys[8] = 100
	keys[9] = 400
	keys[10] = 500
	keys[11] = 900
	keys[12] = 1000

	result := ""
	for i := len(keys) - 1; i >= 0; i-- {
		for num >= keys[i] {
			result += roman[i]
			num -= keys[i]
		}
	}

	return result
}

// A better approach with parallel slices and string builder
func intToRomanAlternative(num int) string {
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	symbols := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var sb strings.Builder

	for i := 0; i < len(values); i++ {
		for num >= values[i] {
			num -= values[i]
			sb.WriteString(symbols[i])
		}
	}

	return sb.String()
}
