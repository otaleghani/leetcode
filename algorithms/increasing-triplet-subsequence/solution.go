package main

import "math"

func increasingTriplet(nums []int) bool {
	i := 0
	j := 1
	k := 2
	freq := make(map[int]int)

	for _, v := range nums {
		freq[v]++
	}

	if len(freq) < 3 {
		return false
	}

	for k < len(nums) {
		if nums[i] >= nums[j] {
			i = j
			j = i + 1
			k = j + 1
			continue
		}

		for kTest := k; kTest < len(nums); kTest++ {
			if nums[kTest] > nums[j] {
				return true
			}
		}

		j = k
		k = k + 1
	}

	return false
}

// A far better alternative using greedy
func increasingTripletAlternative(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first := math.MaxInt32
	second := math.MaxInt32

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}
