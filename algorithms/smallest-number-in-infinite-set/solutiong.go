package main

import "container/heap"

type SmallestInfiniteSetSlow struct {
	nums   []int
	popped map[int]bool
}

func ConstructorSlow() SmallestInfiniteSetSlow {
	nums := make([]int, 1000)
	popped := make(map[int]bool)
	for i := 1; i <= 1000; i++ {
		nums[i-1] = i
		popped[i] = false
	}

	return SmallestInfiniteSetSlow{
		nums:   nums,
		popped: popped,
	}
}

func (this *SmallestInfiniteSetSlow) PopSmallestSlow() int {
	smallest := this.nums[0]
	this.nums = this.nums[1:]
	this.popped[smallest] = true
	return smallest
}

func (this *SmallestInfiniteSetSlow) AddBackSlow(num int) {
	if this.popped[num] {
		for i := 0; i < len(this.nums); i++ {
			if num > this.nums[i] {
				continue
			}
			before := this.nums[0:i]
			after := append([]int{num}, this.nums[i:]...)
			this.nums = append(before, after...)
			this.popped[num] = false
			break
		}
	}
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */

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

type SmallestInfiniteSet struct {
	currentSmallest int
	addedBackHeap   *IntHeap
	inHeapSet       map[int]bool
}

func Constructor() SmallestInfiniteSet {
	h := &IntHeap{}
	heap.Init(h)

	return SmallestInfiniteSet{
		currentSmallest: 1,
		addedBackHeap:   h,
		inHeapSet:       make(map[int]bool),
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	var result int

	if this.addedBackHeap.Len() > 0 && (*this.addedBackHeap)[0] < this.currentSmallest {
		result = heap.Pop(this.addedBackHeap).(int)
		delete(this.inHeapSet, result)
	} else if this.addedBackHeap.Len() > 0 && (*this.addedBackHeap)[0] == this.currentSmallest {
		result = heap.Pop(this.addedBackHeap).(int)
		delete(this.inHeapSet, result)
		this.currentSmallest++
	} else {
		result = this.currentSmallest
		this.currentSmallest++
	}
	return result
}

func (this *SmallestInfiniteSet) AddBack(num int) {
	if num >= this.currentSmallest {
		return
	}
	if this.inHeapSet[num] {
		return
	}
	heap.Push(this.addedBackHeap, num)
	this.inHeapSet[num] = true
}

/**
 * Your SmallestInfiniteSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.PopSmallest();
 * obj.AddBack(num);
 */
