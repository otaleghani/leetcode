package main

import (
	"container/heap"
	"math"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Swap(x, y int) {
	h[x], h[y] = h[y], h[x]
}

func (h IntHeap) Less(x, y int) bool {
	return h[x] < h[y]
}

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func totalCost(costs []int, k int, candidates int) int64 {
	startHeap := IntHeap{}
	endHeap := IntHeap{}
	heap.Init(&startHeap)
	heap.Init(&endHeap)
	pntStart := 0
	pntEnd := len(costs) - 1

	for i := 0; i < candidates; i++ {
		if pntStart > pntEnd {
			break
		}
		heap.Push(&startHeap, costs[pntStart])
		pntStart++

		if pntStart > pntEnd {
			break
		}
		heap.Push(&endHeap, costs[pntEnd])
		pntEnd--
	}

	var total int64 = 0
	for i := 0; i < k; i++ {
		if startHeap.Len() == 0 && endHeap.Len() == 0 {
			break
		}
		startMin := math.MaxInt
		if startHeap.Len() > 0 {
			startMin = startHeap[0]
		}

		endMin := math.MaxInt
		if endHeap.Len() > 0 {
			endMin = endHeap[0]
		}

		if startMin <= endMin {
			cost := heap.Pop(&startHeap).(int)
			total += int64(cost)
			if pntStart <= pntEnd {
				heap.Push(&startHeap, costs[pntStart])
				pntStart++
			}
		} else {
			cost := heap.Pop(&endHeap).(int)
			total += int64(cost)
			if pntStart <= pntEnd {
				heap.Push(&endHeap, costs[pntEnd])
				pntEnd--
			}
		}
	}

	return total
}
