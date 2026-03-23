package main

func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	dp := make([]int, n)

	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}

	for row := n - 2; row >= 0; row-- {
		for i := 0; i <= row; i++ {
			dp[i] = triangle[row][i] + min(dp[i], dp[i+1])
		}
	}

	return dp[0]
}
