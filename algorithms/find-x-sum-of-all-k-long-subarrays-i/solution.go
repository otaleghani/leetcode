package main

import "sort"

func findXSum(nums []int, k int, x int) []int {
	n := len(nums)
	ans := make([]int, 0, n-k+1)

	for i := 0; i <= n-k; i++ {
		freq := make(map[int]int)
		for j := i; j < i+k; j++ {
			freq[nums[j]]++
		}

		type pair struct {
			val   int
			count int
		}
		pairs := make([]pair, 0, len(freq))
		for v, c := range freq {
			pairs = append(pairs, pair{val: v, count: c})
		}

		sort.Slice(pairs, func(a, b int) bool {
			if pairs[a].count == pairs[b].count {
				return pairs[a].val > pairs[b].val
			}
			return pairs[a].count > pairs[b].count
		})

		currentSum := 0
		limit := x
		if len(pairs) < x {
			limit = len(pairs)
		}

		for p := 0; p < limit; p++ {
			currentSum += pairs[p].val * pairs[p].count
		}

		ans = append(ans, currentSum)
	}

	return ans
}
