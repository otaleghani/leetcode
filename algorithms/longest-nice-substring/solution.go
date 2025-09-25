package main

import "unicode"

func longestNiceSubstring(s string) string {
	if len(s) < 2 {
		return ""
	}

	charSet := make(map[rune]bool)
	for _, char := range s {
		charSet[char] = true
	}

	// Find bad characters, divide and conquer
	for i, char := range s {
		if unicode.IsLower(char) && !charSet[unicode.ToUpper(char)] {
			leftSub := longestNiceSubstring(s[0:i])
			rightSub := longestNiceSubstring(s[i+1:])
			if len(leftSub) >= len(rightSub) {
				return leftSub
			}
			return rightSub
		}
		if unicode.IsUpper(char) && !charSet[unicode.ToLower(char)] {
			leftSub := longestNiceSubstring(s[0:i])
			rightSub := longestNiceSubstring(s[i+1:])
			if len(leftSub) >= len(rightSub) {
				return leftSub
			}
			return rightSub
		}
	}

	return s
}
