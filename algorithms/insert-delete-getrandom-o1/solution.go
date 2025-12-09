package main

import "math/rand"

type RandomizedSet struct {
	nums       []int
	valToIndex map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		nums:       []int{},
		valToIndex: make(map[int]int),
	}
}

func (this *RandomizedSet) Insert(val int) bool {
	if _, exists := this.valToIndex[val]; exists {
		return false
	}

	this.nums = append(this.nums, val)
	this.valToIndex[val] = len(this.nums) - 1

	return true
}

func (this *RandomizedSet) Remove(val int) bool {
	indexToRemove, exists := this.valToIndex[val]
	if !exists {
		return false
	}

	lastIndex := len(this.nums) - 1
	lastVal := this.nums[lastIndex]

	this.nums[indexToRemove] = lastVal
	this.valToIndex[lastVal] = indexToRemove

	this.nums = this.nums[:lastIndex]
	delete(this.valToIndex, val)

	return true
}

func (this *RandomizedSet) GetRandom() int {
	randomIndex := rand.Intn(len(this.nums))
	return this.nums[randomIndex]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
