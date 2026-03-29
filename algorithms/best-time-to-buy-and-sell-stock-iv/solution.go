package main

import "math"

func maxProfit(k int, prices []int) int {
	n := len(prices)
	if n == 0 || k == 0 {
		return 0
	}

	if k >= n/2 {
		maxProf := 0
		for i := 1; i < n; i++ {
			if prices[i] > prices[i-1] {
				maxProf += prices[i] - prices[i-1]
			}
		}
		return maxProf
	}

	buy := make([]int, k+1)
	sell := make([]int, k+1)

	for i := 0; i <= k; i++ {
		buy[i] = math.MinInt32
	}

	for _, price := range prices {
		for i := 1; i <= k; i++ {
			buy[i] = max(buy[i], sell[i-1]-price)

			sell[i] = max(sell[i], buy[i]+price)
		}
	}

	return sell[k]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
