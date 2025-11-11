package main

type SmallestInfiniteSet struct {
	nums   []int
	popped map[int]bool
}

func Constructor() SmallestInfiniteSet {
	nums := make([]int, 1000)
	popped := make(map[int]bool)
	for i := 1; i <= 1000; i++ {
		nums[i-1] = i
		popped[i] = false
	}

	return SmallestInfiniteSet{
		nums:   nums,
		popped: popped,
	}
}

func (this *SmallestInfiniteSet) PopSmallest() int {
	smallest := this.nums[0]
	this.nums = this.nums[1:]
	this.popped[smallest] = true
	return smallest
}

func (this *SmallestInfiniteSet) AddBack(num int) {
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
