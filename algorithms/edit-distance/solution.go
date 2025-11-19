package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1))
	for i := range dp {
		dp[i] = make([]int, len(word2))
		for j := range dp[i] {
			dp[i][j] = 501
		}
	}

	var f func(x, y int) int
	f = func(x, y int) int {
		if x == len(word1) {
			return len(word2) - y
		}
		if y == len(word2) {
			return len(word1) - x
		}

		if dp[x][y] != 501 {
			return dp[x][y]
		}

		// Is the current the same?
		if word1[x] == word2[y] {
			dp[x][y] = f(x+1, y+1)
			return dp[x][y]
		}

		// If is not the same...
		insert := 1 + f(x, y+1)
		remove := 1 + f(x+1, y)
		replace := 1 + f(x+1, y+1)

		dp[x][y] = min(insert, remove, replace)
		return dp[x][y]
	}

	return f(0, 0)
}

func min(x, y, z int) int {
	if x <= y && x <= z {
		return x
	} else if y <= x && y <= z {
		return y
	} else {
		return z
	}
}
