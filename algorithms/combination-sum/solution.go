package main

func combinationSum(candidates []int, target int) [][]int {
	var res [][]int
	var current []int

	var dfs func(start int, currentSum int)
	dfs = func(start int, currentSum int) {
		if currentSum == target {
			temp := make([]int, len(current))
			copy(temp, current)
			res = append(res, temp)
			return
		}

		if currentSum > target {
			return
		}

		for i := start; i < len(candidates); i++ {
			current = append(current, candidates[i])
			dfs(i, currentSum+candidates[i])
			current = current[:len(current)-1]
		}
	}

	dfs(0, 0)
	return res
}
