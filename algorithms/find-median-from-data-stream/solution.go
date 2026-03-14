package main

import "container/heap"

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// MedianFinder Implementation
type MedianFinder struct {
	minH *MinHeap
	maxH *MaxHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minH: &MinHeap{},
		maxH: &MaxHeap{},
	}
}

func (this *MedianFinder) AddNum(num int) {
	heap.Push(this.maxH, num)
	heap.Push(this.minH, heap.Pop(this.maxH))

	if this.minH.Len() > this.maxH.Len() {
		heap.Push(this.maxH, heap.Pop(this.minH))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.maxH.Len() > this.minH.Len() {
		return float64((*this.maxH)[0])
	}

	return (float64((*this.maxH)[0]) + float64((*this.minH)[0])) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
