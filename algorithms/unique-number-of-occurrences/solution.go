package main

func uniqueOccurrences(arr []int) bool {
	freq := make(map[int]int)
	for _, v := range arr {
		freq[v]++
	}

	otherMap := make(map[int]int)
	for _, v := range freq {
		otherMap[v]++
	}

	return len(freq) == len(otherMap)
}

// A more idiomatic way, uses less memory too!
func uniqueOccurrencesAlternative(arr []int) bool {
	freqCount := make(map[int]int)
	for _, v := range arr {
		freqCount[v]++
	}

	freqSeen := make(map[int]struct{})
	for _, v := range freqCount {
		if _, seen := freqSeen[v]; seen {
			return false
		}
		freqSeen[v] = struct{}{}
	}

	return true
}
