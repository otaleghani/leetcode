package main

type Key struct {
	i int
	m int
	n int
}

// Cache is really heavy on the memory side
func findMaxFormSlow(strs []string, m int, n int) int {
	cache := make(map[Key]int)

	var dfs func(i, m, n int) int
	dfs = func(i, m, n int) int {
		if i == len(strs) {
			return 0
		}
		key := Key{i: i, m: m, n: n}
		if val, ok := cache[key]; ok {
			return val
		}

		mCount, nCount := count(strs[i])

		res := dfs(i+1, m, n) // option 1: skip one

		if mCount <= m && nCount <= n {
			res = max(res, 1+dfs(i+1, m-mCount, n-nCount))
		}

		cache[key] = res
		return res
	}

	return dfs(0, m, n)
}

func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeros, ones := count(str)
		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				dp[i][j] = max(dp[i][j], 1+dp[i-zeros][j-ones])
			}
		}
	}

	return dp[m][n]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func count(str string) (m, n int) {
	for _, v := range str {
		if v == '0' {
			m += 1
		}
		if v == '1' {
			n += 1
		}
	}
	return
}
