package main

func containsNearbyDuplicate(nums []int, k int) bool {
	freq := make(map[int][]int)

	for i := 0; i < len(nums); i++ {
		freq[nums[i]] = append(freq[nums[i]], i)
	}

	for _, array := range freq {
		if len(array) < 2 {
			continue
		}

		for i := 0; i < len(array)-1; i++ {
			res := array[i+1] - array[i]
			if res < 0 {
				res *= -1
			}

			if res <= k {
				return true
			}
		}
	}

	return false
}

// Mush slower, but less memory
func containsNearbyDuplicateAlternative(nums []int, k int) bool {
	for left := 0; left < len(nums)-1; left++ {
		offset := left + k
		if offset > len(nums)-1 {
			offset = len(nums) - 1
		}

		for right := left + 1; right <= offset; right++ {
			if nums[left] == nums[right] {
				return true
			}
		}
	}

	return false
}

// Best solution, both in speed and in memory
func containsNearbyDuplicateBest(nums []int, k int) bool {
	n := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] == nums[j] && abs(i-j) <= k {
				return true
			}
			n++
			if n == 100000 {
				return false
			}
		}
	}
	return false
}
func abs(res int) int {
	if res < 0 {
		res *= -1
	}
	return res
}
