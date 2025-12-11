package main

func romanToInt(s string) int {
	conv := make(map[byte]int)
	conv['I'] = 1
	conv['V'] = 5
	conv['X'] = 10
	conv['L'] = 50
	conv['C'] = 100
	conv['D'] = 500
	conv['M'] = 1000

	result := 0
	for i := 0; i < len(s)-1; i++ {
		currVal := conv[s[i]]
		nextVal := conv[s[i+1]]

		if currVal < nextVal {
			result -= currVal
		} else {
			result += currVal
		}
	}

	result += conv[s[len(s)-1]]

	return result
}
