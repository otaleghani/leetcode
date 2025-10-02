package main

func findRepeatedDnaSequences(s string) []string {
	var result []string
	freq := make(map[string]int)

	for i := 9; i < len(s); i++ {
		freq[string(s[i-9:i+1])] += 1
	}

	for i, v := range freq {
		if v > 1 {
			result = append(result, i)
		}
	}

	return result
}

// Single pass, slightly more efficient
func findRepeatedDnaSequencesAlternative(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	var result []string
	freq := make(map[string]int)

	for i := 9; i < len(s); i++ {
		sequence := s[i-9 : i+1]
		freq[sequence] += 1
		if freq[sequence] == 2 {
			result = append(result, sequence)
		}
	}

	return result
}

// Rolling hash
func findRepeatedDnaSequencesRollingHash(s string) []string {
	if len(s) < 10 {
		return []string{}
	}

	charToBits := map[byte]int{
		'A': 0,
		'C': 1,
		'G': 2,
		'T': 3,
	}

	freq := make(map[int]int)
	var result []string
	mask := (1 << 20) - 1
	currentHash := 0

	for i := 0; i < 9; i++ {
		currentHash = (currentHash << 2) | charToBits[s[i]]
	}

	for i := 9; i < len(s); i++ {
		// Calculate hash
		currentHash = ((currentHash << 2) | charToBits[s[i]]) & mask
		freq[currentHash]++
		if freq[currentHash] == 2 {
			result = append(result, s[i-9:i+1])
		}
	}

	return result
}
