package main

import (
	"math/rand"
	"sort"
)

type Solution struct {
	prefixSums []int
}

func Constructor(w []int) Solution {
	prefixSums := make([]int, len(w))
	sum := 0
	for i, weight := range w {
		sum += weight
		prefixSums[i] = sum
	}
	return Solution{prefixSums: prefixSums}
}

func (this *Solution) PickIndex() int {
	target := rand.Intn(this.prefixSums[len(this.prefixSums)-1]) + 1
	return sort.SearchInts(this.prefixSums, target)
}
