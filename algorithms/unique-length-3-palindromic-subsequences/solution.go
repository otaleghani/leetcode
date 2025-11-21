package main

func countPalindromicSubsequence(s string) int {
	first := make([]int, 26)
	last := make([]int, 26)

	for i := range first {
		first[i] = -1
		last[i] = -1
	}

	for i, char := range s {
		idx := char - 'a'
		if first[idx] == -1 {
			first[idx] = i
		}
		last[idx] = i
	}

	result := 0
	for i := 0; i < 26; i++ {
		start := first[i]
		end := last[i]

		if start == -1 || end <= start+1 {
			continue
		}

		uniqueMiddles := make(map[byte]bool)
		for k := start + 1; k < end; k++ {
			uniqueMiddles[s[k]] = true
		}

		result += len(uniqueMiddles)
	}

	return result
}
