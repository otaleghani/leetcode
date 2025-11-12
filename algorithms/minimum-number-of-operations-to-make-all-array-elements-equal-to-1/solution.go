package main

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func minOperations(nums []int) int {
	n := len(nums)
	ones := 0

	for _, x := range nums {
		if x == 1 {
			ones++
		}
	}

	if ones > 0 {
		return n - ones
	}

	minL := -1

	for l := 2; l <= n; l++ { // Iterate through all possible lengths L
		for i := 0; i <= n-l; i++ { // Iterate through all start indices i

			g := 0
			for j := i; j < i+l; j++ {
				if j == i {
					g = nums[j]
				} else {
					g = gcd(g, nums[j])
				}
			}

			if g == 1 {
				minL = l
				break
			}
		}
		if minL != -1 {
			break
		}
	}

	if minL == -1 {
		return -1
	}

	return (minL - 1) + (n - 1)
}
