package main

import "math"

func reverse(x int) int {
	result := 0

	for x != 0 {
		pop := x % 10
		x /= 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && pop > 7) {
			return 0
		}

		if result < math.MinInt32/10 || (result == math.MinInt32/10 && pop < -8) {
			return 0
		}

		result = result*10 + pop
	}

	return result
}
