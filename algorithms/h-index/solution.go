package main

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)

	count := 0 // Minimum citations
	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] <= count {
			break
		}
		if citations[i] != 0 {
			count += 1
		}
	}

	return count
}
