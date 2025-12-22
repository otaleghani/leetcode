package main

import "strings"

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")
	wordMap := make(map[byte]string)
	patternMap := make(map[string]struct{})

	if len(words) != len(pattern) {
		return false
	}

	for idx, currWord := range words {
		key := pattern[idx]

		savedWord, ok := wordMap[key]
		if !ok {
			wordMap[key] = currWord
			// Was the pattern already used?
			_, ok := patternMap[currWord]
			if !ok { // Pattern never used
				patternMap[currWord] = struct{}{}
			} else {
				return false
			}
			continue
		}
		if savedWord != currWord {
			return false
		}
	}

	return true
}
