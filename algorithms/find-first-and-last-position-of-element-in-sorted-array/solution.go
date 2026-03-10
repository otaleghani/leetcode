package main

func searchRange(nums []int, target int) []int {
	// Find the target
	left := 0
	right := len(nums) - 1
	start := -1
	end := -1

	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			start = mid
			end = mid
			for start >= 0 && nums[start] == target {
				start--
			}
			start++

			for end <= len(nums)-1 && nums[end] == target {
				end++
			}
			end--
			break
		}

		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return []int{start, end}
}
