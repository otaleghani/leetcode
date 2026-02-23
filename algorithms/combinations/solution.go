package main

func combine(n int, k int) [][]int {
	comb := [][]int{}
	curr := []int{}

	var f func(start int)
	f = func(start int) {
		if len(curr) == k {
			dCopy := make([]int, k)
			copy(dCopy, curr)
			comb = append(comb, dCopy)
		}

		for i := start; i <= n; i++ {
			curr = append(curr, i)
			f(i + 1)
			curr = curr[:len(curr)-1]
		}
	}

	f(1)
	return comb
}
