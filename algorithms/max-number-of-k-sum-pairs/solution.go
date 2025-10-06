package main

func maxOperations(nums []int, k int) int {
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	ops := 0
	for i := range freq {
		if _, ok := freq[k-i]; ok {
			for {
				if i == k-i {
					if freq[i] >= 2 {
						freq[i] -= 2
						ops++
						continue
					} else {
						break
					}
				}
				if freq[i] > 0 && freq[k-i] > 0 {
					freq[i]--
					freq[k-i]--
					ops++
				} else {
					break
				}
			}
		}
	}
	return ops
}

// This should be faster, but its worst in both speed and memory. idk why
func maxOperationsAlternative(nums []int, k int) int {
	freq := make(map[int]int)
	ops := 0

	for _, v := range nums {
		freq[v]++
	}

	for num, count := range freq {
		if count == 0 {
			continue
		}

		complement := k - num

		if num == complement {
			ops += count / 2
			continue
		}

		if compCount, ok := freq[complement]; ok && compCount > 0 {
			pairs := min(count, compCount)
			ops += pairs
			freq[num] = 0
			freq[complement] = 0
		}
	}

	return ops
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
