package main

import (
	"container/heap"
	"sort"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}

func (h MaxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *MaxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

type Project struct {
	profit  int
	capital int
}

func findMaximizedCapital(k int, w int, profits []int, capital []int) int {
	projects := make([]Project, len(profits))

	for i := 0; i < len(profits); i++ {
		projects[i] = Project{profit: profits[i], capital: capital[i]}
	}

	sort.Slice(projects, func(i, j int) bool {
		return projects[i].capital < projects[j].capital
	})

	h := &MaxHeap{}
	heap.Init(h)

	ptr := 0

	for i := 0; i < k; i++ {
		for ptr < len(profits) && projects[ptr].capital <= w {
			heap.Push(h, projects[ptr].profit)
			ptr++
		}

		if h.Len() == 0 {
			break
		}

		w += heap.Pop(h).(int)
	}

	return w
}
