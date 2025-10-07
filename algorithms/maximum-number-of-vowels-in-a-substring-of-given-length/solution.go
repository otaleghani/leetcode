package main

func maxVowels(s string, k int) int {
	left := 0
	count := 0
	max := 0
	for right := 0; right < len(s); right++ {
		if isVowel(s[right]) {
			count++

		}

		if right >= k {
			if isVowel(s[left]) {
				count--
			}
			left++
		}
		if count > max {
			max = count
		}
	}
	return max
}

func isVowel(b byte) bool {
	switch b {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	default:
		return false
	}
}
