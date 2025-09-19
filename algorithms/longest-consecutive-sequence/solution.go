package main

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	// Convert the array into a hash map
	hashMap := make(map[int]int)

	for _, v := range nums {
		hashMap[v] = 1
	}

	max := 1
	for i, _ := range hashMap {
		// Find the start of the array which is if record before this one exists
		if hashMap[i-1] == 1 {
			continue
		} else {
			currentMax := 1
			j := 1
			for {
				if hashMap[i+j] == 1 {
					currentMax = currentMax + 1
					j = j + 1
				} else {
					break
				}
			}
			if currentMax > max {
				max = currentMax
			}
		}
	}

	return max
}
