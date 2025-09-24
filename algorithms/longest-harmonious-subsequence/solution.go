package main

func findLHS(nums []int) int {
	freq := make(map[int]int)
	for _, n := range nums {
		freq[n]++
	}

	result := 0

	for k, v := range freq {
		if next, ok := freq[k+1]; ok {
			if v+next > result {
				result = v + next
			}
		}
	}

	return result
}
