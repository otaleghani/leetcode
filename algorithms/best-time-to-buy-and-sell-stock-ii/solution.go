package main

func maxProfit(prices []int) int {
	var f func(day int) (int, int)
	f = func(day int) (int, int) {
		if day == 0 {
			return -prices[day], 0
		}

		prevHold, prevNotHold := f(day - 1)

		hold := max(prevHold, prevNotHold-prices[day])
		notHold := max(prevNotHold, prevHold+prices[day])

		return hold, notHold
	}

	lastDay := len(prices) - 1
	_, finalNotHold := f(lastDay)
	return finalNotHold
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
