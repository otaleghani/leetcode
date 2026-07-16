package main

import "sort"

func numSubseq(nums []int, target int) int {
	mod := 1000000007
	sort.Ints(nums)
	n := len(nums)

	pows := make([]int, n)
	pows[0] = 1
	for i := 1; i < n; i++ {
		pows[i] = (pows[i-1] * 2) % mod
	}

	ans := 0
	l, r := 0, n-1

	for l <= r {
		if nums[l]+nums[r] > target {
			r--
		} else {
			ans = (ans + pows[r-l]) % mod
			l++
		}
	}

	return ans
}
