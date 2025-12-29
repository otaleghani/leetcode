package main

func isHappy(n int) bool {
	encounters := make(map[int]struct{})

	for n != 0 {
		digits := []int{}
		for n > 0 {
			digits = append(digits, n%10)
			n /= 10
		}

		n = 0
		for _, val := range digits {
			n += (val * val)
		}

		if n == 1 {
			return true
		}

		if _, exists := encounters[n]; exists {
			return false
		}

		encounters[n] = struct{}{}
	}

	return true
}
