package main

func characterReplacement(s string, k int) int {
	charSet := make(map[byte]bool)
	for i := range s {
		charSet[s[i]] = true
	}
	result := 0
	for c := range charSet {
		count, end := 0, 0
		for start := 0; start < len(s); start++ {
			if s[start] == c {
				count++
			}

			for (start-end+1)-count > k {
				if s[end] == c {
					count--
				}
				end++
			}

			result = max(result, start-end+1)
		}
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
