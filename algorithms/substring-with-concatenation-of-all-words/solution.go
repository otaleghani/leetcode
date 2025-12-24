package main

// This solution is slow but was the first one that came to mind first thing in the morning
func findSubstring(s string, words []string) []int {
	result := []int{}
	lenWords := len(words[0])
	concatLen := len(words[0]) * len(words)
	wordsMap := make(map[string]struct{})
	for _, word := range words {
		wordsMap[word] = struct{}{}
	}

	for i := 0; i < len(s); i++ {
		if i+lenWords > len(s) {
			break
		}
		word := s[i : i+lenWords]
		_, ok := wordsMap[word]
		if !ok { // Word does not exist, continue
			continue
		} else {
			if i+concatLen > len(s) {
				break
			}
			testString := s[i : i+concatLen]
			if isPermutation(testString, words, lenWords) {
				result = append(result, i)
			}
		}
	}

	return result
}

func isPermutation(s string, words []string, lenWords int) bool {
	wordsToUse := make(map[string]int)
	for _, word := range words {
		wordsToUse[word] += 1
	}

	i := 0
	for i < len(s) {
		wordToTest := s[i : i+lenWords]
		_, ok := wordsToUse[wordToTest]
		if ok {
			wordsToUse[wordToTest] -= 1
			if wordsToUse[wordToTest] == 0 {
				delete(wordsToUse, wordToTest)
			}
			i += lenWords
		} else {
			return false
		}
	}

	return true
}

// Better sliding window approach
func findSubstringAlternative(s string, words []string) []int {
	if len(s) == 0 || len(words) == 0 {
		return []int{}
	}

	result := []int{}
	wordLen := len(words[0])
	numWords := len(words)
	sLen := len(s)

	expectedCounts := make(map[string]int)
	for _, word := range words {
		expectedCounts[word]++
	}

	for i := 0; i < wordLen; i++ {
		left := i
		right := i
		currentCounts := make(map[string]int)
		count := 0 // Number of valid words currently in the window

		// Slide the window across s
		for right+wordLen <= sLen {
			w := s[right : right+wordLen]
			right += wordLen

			if _, ok := expectedCounts[w]; !ok {
				currentCounts = make(map[string]int)
				count = 0
				left = right
			} else {
				currentCounts[w]++
				count++

				for currentCounts[w] > expectedCounts[w] {
					leftWord := s[left : left+wordLen]
					currentCounts[leftWord]--
					count--
					left += wordLen
				}

				if count == numWords {
					result = append(result, left)
					firstWord := s[left : left+wordLen]
					currentCounts[firstWord]--
					count--
					left += wordLen
				}
			}
		}
	}

	return result
}
