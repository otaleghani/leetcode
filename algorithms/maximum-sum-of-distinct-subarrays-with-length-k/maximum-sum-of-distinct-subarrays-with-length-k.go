package main

import "fmt"

func main() {
  array := []int{1,5,4,2,9,9,9}
  fmt.Print(maximumSubarraySum(array, 3))
}

// Failed attempt
func maximumSubarraySum(nums []int, k int) int64 {
    result := 0

    for i := 0; i < len(nums)-k+1; i++ {
        m := make(map[int]int)
        for j := i; j < i+k; j++ {
            m[nums[j]] = j
        }
        if len(m) == k {
            fmt.Print(m)
            newResult := 0
            for ind, _ := range m {
                newResult = newResult + ind
            }
            if result < newResult {
                result = newResult
            }
        }
    }
    return int64(result)
}

// Sliding windows
func maximumSubarraySum(nums []int, k int) int64 {
	if k > len(nums) {
		return 0
	}

	// Map to store the frequency of elements in the current window
	countMap := make(map[int]int)
	var maxSum int64 = 0
	var currentSum int64 = 0
	valid := false

	// Initialize the first window
	for i := 0; i < k; i++ {
		countMap[nums[i]]++
		currentSum += int64(nums[i])
	}

	// Check if the first window is valid
	if len(countMap) == k {
		maxSum = currentSum
		valid = true
	}

	// Slide the window
	for i := k; i < len(nums); i++ {
		// Remove the element exiting the window
		exiting := nums[i-k]
		countMap[exiting]--
		if countMap[exiting] == 0 {
			delete(countMap, exiting)
		}
		currentSum -= int64(exiting)

		// Add the new element entering the window
		entering := nums[i]
		countMap[entering]++
		currentSum += int64(entering)

		// Check if the current window has all distinct elements
		if len(countMap) == k {
			if !valid || currentSum > maxSum {
				maxSum = currentSum
				valid = true
			}
		}
	}

	if valid {
		return maxSum
	}
	return 0
}

