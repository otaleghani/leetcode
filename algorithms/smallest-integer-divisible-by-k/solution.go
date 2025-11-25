package main

func smallestRepunitDivByK(k int) int {
	remainder := 0

	for i := 1; i < 100000; i++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return i
		}
	}

	return -1
}
