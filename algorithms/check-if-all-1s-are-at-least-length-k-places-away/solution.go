package main

func kLengthApart(nums []int, k int) bool {
	last := 0
	notFirst := false

	for i, v := range nums {
		if v == 1 {
			if i-last-1 < k && notFirst {
				return false
			}
			last = i
			notFirst = true
		}

	}

	return true
}
