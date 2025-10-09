package main

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	wordOneMap := make(map[byte]int)
	wordTwoMap := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		wordOneMap[word1[i]]++
		wordTwoMap[word2[i]]++
	}

	wordOneFreqMap := make(map[int]int)
	wordTwoFreqMap := make(map[int]int)

	for _, v := range wordOneMap {
		wordOneFreqMap[v]++
	}
	for _, v := range wordTwoMap {
		wordTwoFreqMap[v]++
	}

	for idxOne, _ := range wordOneMap {
		_, ok := wordTwoMap[idxOne]
		if !ok {
			return false
		}
	}
	for idxOne, valOne := range wordOneFreqMap {
		if valTwo, ok := wordTwoFreqMap[idxOne]; ok {
			if valOne != valTwo {
				return false
			}
		} else {
			return false
		}
	}

	return true
}

// A more idiomatic way of solving this
func closeStringsAlternative(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}

	freq1 := make([]int, 26)
	freq2 := make([]int, 26)
	for i := 0; i < len(word1); i++ {
		freq1[word1[i]-'a']++
		freq2[word2[i]-'a']++
	}

	// Condition one
	for i := 0; i < 26; i++ {
		if (freq1[i] > 0 && freq2[i] == 0) || (freq1[i] == 0 && freq2[i] > 0) {
			return false
		}
	}

	// Condition two
	sort.Ints(freq1)
	sort.Ints(freq2)

	for i := 0; i < 26; i++ {
		if freq1[i] != freq2[i] {
			return false
		}
	}

	return true
}
