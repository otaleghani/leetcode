package main

// This solution fails for time limit
func maxProfit(prices []int, fee int) int {
	var f func(x, y int) int
	f = func(x, y int) int {
		if x >= len(prices) || y >= len(prices) {
			return 0
		}
		if x < y {
			currentTrade := prices[y] - prices[x] - fee
			sell := currentTrade + f(y+1, y+1)
			hold := f(x, y+1)
			return max(sell, hold)
		}

		buy := 0
		skip := f(y+1, y+1)
		if x == y {
			buy = f(x, y+1)
		}

		return max(buy, skip)
	}

	return f(0, 0)
}

// This doesn't fail and works
func maxProfitAlternative(prices []int, fee int) int {
	dp := make([][]*int, len(prices))
	for i := range dp {
		dp[i] = make([]*int, 2)
	}

	var f func(i int, holding int) int
	f = func(i int, holding int) int {
		if i >= len(prices) {
			return 0
		}

		if dp[i][holding] != nil {
			return *dp[i][holding]
		}

		res := 0
		if holding == 1 {
			sell := f(i+1, 0) + prices[i] - fee
			hold := f(i+1, 1)
			res = max(sell, hold)
		} else {
			buy := f(i+1, 1) - prices[i]
			skip := f(i+1, 0)
			res = max(buy, skip)
		}

		dp[i][holding] = &res
		return res
	}

	return f(0, 0)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
