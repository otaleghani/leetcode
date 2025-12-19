package main

func threeSum(nums []int) [][]int {
	result := make([][]int, 0)

	sort.Ints(nums) // Sort the input array

	for i := 0; i < len(nums)-2; i++ {
		// Skip duplicates
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left := i + 1
		right := len(nums) - 1

		for left < right {
			currSum := nums[i] + nums[left] + nums[right]

			if currSum == 0 {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				// Skip duplicates
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if currSum < 0 {
				left++
			} else {
				right--
			}
		}
	}

	return result
}
