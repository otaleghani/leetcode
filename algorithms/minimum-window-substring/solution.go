package main

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	tFreq := make(map[byte]int)
	for i := range t {
		tFreq[t[i]]++
	}

	left := 0
	resStart, resLen := -1, len(s)+1

	have, need := 0, len(tFreq)
	windowFreq := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		char := s[right]
		windowFreq[char]++

		if _, ok := tFreq[char]; ok && windowFreq[char] == tFreq[char] {
			have++
		}

		for have == need {
			currentLen := right - left + 1
			if currentLen < resLen {
				resLen = currentLen
				resStart = left
			}

			leftChar := s[left]
			windowFreq[leftChar]--

			if _, ok := tFreq[leftChar]; ok && windowFreq[leftChar] < tFreq[leftChar] {
				have--
			}

			left++
		}
	}

	if resStart == -1 {
		return ""
	}

	return s[resStart : resStart+resLen]
}
