package main

func countOdds(low int, high int) int {
	test := high - low + 1
	if test%2 == 0 {
		return test / 2
	}
	if high%2 == 0 || low%2 == 0 {
		return test / 2
	}
	return (test / 2) + 1
}
