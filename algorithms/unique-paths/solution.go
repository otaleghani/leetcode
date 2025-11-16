package main

func uniquePaths(m int, n int) int {

	dp := make([][]int, m)
	for i := range m {
		dp[i] = make([]int, n)
	}

	var f func(x, y, count int) int
	f = func(x, y, count int) int {
		if x == m-1 || y == n-1 {
			return count + 1
		}

		if dp[x][y] != 0 {
			return dp[x][y]
		}

		right := f(x+1, y, count)
		down := f(x, y+1, count)

		dp[x][y] = right + down
		return dp[x][y]
	}

	return f(0, 0, 0)
}
