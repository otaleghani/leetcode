package main

import "container/heap"

type Item struct {
	sum int
	i   int
	j   int
}

type MinHeap []Item

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].sum < h[j].sum }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(Item))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kSmallestPairs(nums1 []int, nums2 []int, k int) [][]int {
	h := &MinHeap{}
	heap.Init(h)

	for i := 0; i < len(nums1) && i < k; i++ {
		heap.Push(h, Item{
			sum: nums1[i] + nums2[0],
			i:   i,
			j:   0,
		})
	}

	var res [][]int

	for h.Len() > 0 && k > 0 {
		item := heap.Pop(h).(Item)
		res = append(res, []int{nums1[item.i], nums2[item.j]})
		k--

		if item.j+1 < len(nums2) {
			heap.Push(h, Item{
				sum: nums1[item.i] + nums2[item.j+1],
				i:   item.i,
				j:   item.j + 1,
			})
		}
	}

	return res
}
