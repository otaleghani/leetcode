package main

import (
	"container/heap"
	"sort"
)

type pair struct {
	a int // nums1
	b int // nums2
}

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func maxScore(nums1 []int, nums2 []int, k int) int64 {
	// Sort based on nums2 in ascending order
	pairs := make([]pair, len(nums1))
	for i := 0; i < len(nums1); i++ {
		pairs[i] = pair{a: nums1[i], b: nums2[i]}
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].b > pairs[j].b
	})

	minHeap := &IntHeap{}
	heap.Init(minHeap)
	var currentSum int64 = 0
	var max int64

	for i := 0; i < len(pairs); i++ {
		heap.Push(minHeap, pairs[i].a)
		currentSum += int64(pairs[i].a)

		if minHeap.Len() > k {
			smallest := heap.Pop(minHeap).(int)
			currentSum -= int64(smallest)
		}

		if minHeap.Len() == k {
			score := currentSum * int64(pairs[i].b)
			if score > max {
				max = score
			}
		}
	}

	return max
}
