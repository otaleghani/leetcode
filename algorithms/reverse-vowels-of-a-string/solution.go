package main

func reverseVowels(s string) string {
	sRune := []rune(s)
	left := len(sRune) - 1

cool:
	for right := 0; right < left; right++ {
		if isVowel(sRune[right]) {
			for {
				if left <= right {
					break cool
				}
				if isVowel(sRune[left]) {
					break
				}
				left--
			}
			current := sRune[right]
			sRune[right] = sRune[left]
			sRune[left] = current
			left--
		}
	}
	return string(sRune)
}

func isVowel(c rune) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	}
	return false
}

// A more coincise solution
func reverseVowelsAlternative(s string) string {
	sRune := []rune(s)
	left, right := 0, len(sRune)-1

	for left < right {
		if !isVowel(sRune[left]) {
			left++
			continue
		}
		if !isVowel(sRune[right]) {
			right--
			continue
		}

		sRune[left], sRune[right] = sRune[right], sRune[left]
		left++
		right--
	}

	return string(sRune)
}
