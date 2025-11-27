package main

func findErrorNums(nums []int) []int {
	dup := 0
	miss := 0
	freq := make(map[int]int)
	for _, v := range nums {
		freq[v] += 1
	}
	for i := 1; i <= len(nums); i++ {
		if freq[i] == 0 {
			miss = i
		} else if freq[i] == 2 {
			dup = i
		}
	}

	return []int{dup, miss}
}
