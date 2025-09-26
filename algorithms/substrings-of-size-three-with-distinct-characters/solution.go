package main

func countGoodSubstrings(s string) int {
	// Create a current hashmap
	left := 0
	result := 0
	freq := make(map[byte]int)
	for right := 0; right < len(s); right++ {
		freq[s[right]] += 1

		if right-left == 2 {
			// check from left to right
			if len(freq) == 3 {
				result += 1
			}

			freq[s[left]] -= 1
			if freq[s[left]] == 0 {
				delete(freq, s[left])
			}
			left++
		}
	}
	return result
}
