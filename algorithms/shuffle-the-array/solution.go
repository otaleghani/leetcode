package main

func shuffle(nums []int, n int) []int {
	pnt1 := 0
	pnt2 := n
	newArr := make([]int, len(nums))
	for i := range nums {
		if i%2 == 0 {
			newArr[i] = nums[pnt1]
			pnt1++
		} else {
			newArr[i] = nums[pnt2]
			pnt2++
		}
	}

	return newArr
}
