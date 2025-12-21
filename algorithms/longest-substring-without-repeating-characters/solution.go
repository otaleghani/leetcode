package main

func lengthOfLongestSubstring(s string) int {
	pos := make(map[byte]int)
	left := 0
	max := 0
	for right := 0; right < len(s); right++ {
		char := s[right]
		_, present := pos[char]
		if present && pos[char] >= left {
			left = pos[char] + 1
		}
		pos[char] = right
		currMax := right - left + 1
		if currMax > max {
			max = currMax
		}
	}

	return max
}
