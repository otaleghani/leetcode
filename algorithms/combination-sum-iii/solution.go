package main

func combinationSum3(k int, n int) [][]int {
	var results [][]int
	var currentPath []int

	var backtrack func(start int, currentSum int)
	backtrack = func(start int, currentSum int) {
		if currentSum > n || len(currentPath) > k {
			return
		}

		if len(currentPath) == k && currentSum == n {
			temp := make([]int, k)
			copy(temp, currentPath)
			results = append(results, temp)
			return
		}

		for i := start; i <= 9; i++ {
			currentPath = append(currentPath, i)
			backtrack(i+1, currentSum+i)
			currentPath = currentPath[:len(currentPath)-1]
		}
	}

	backtrack(1, 0)
	return results
}
