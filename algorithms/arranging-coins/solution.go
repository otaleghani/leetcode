package main

func arrangeCoins(n int) int {
	left, right := 1, n
	for left <= right {
		mid := left + (right-left)/2
		coins := (mid * (mid + 1)) / 2

		if coins == n {
			return mid
		} else if coins < n {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return right
}
