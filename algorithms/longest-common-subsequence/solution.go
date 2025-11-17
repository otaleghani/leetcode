package main

func longestCommonSubsequence(text1 string, text2 string) int {
	dp := make([][]int, len(text1)+1)
	for i := range dp {
		dp[i] = make([]int, len(text2)+1)
	}
	for i := range dp {
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(x, y int) int
	f = func(x, y int) int {
		if x == len(dp)-1 || y == len(dp[0])-1 {
			return 0
		}

		if dp[x][y] != -1 {
			return dp[x][y]
		}

		if text1[x] == text2[y] {
			return 1 + f(x+1, y+1)
		}

		left := f(x+1, y)
		down := f(x, y+1)

		dp[x][y] = max(left, down)
		return dp[x][y]
	}

	return f(0, 0)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
