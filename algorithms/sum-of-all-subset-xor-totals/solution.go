package main

func subsetXORSum(nums []int) int {
	return backtrack(0, 0, nums)
}

func backtrack(index, currentXOR int, nums []int) int {
	if index == len(nums) {
		return currentXOR
	}
	xorSum_without_element := backtrack(index+1, currentXOR, nums)
	xorSum_with_element := backtrack(index+1, currentXOR^nums[index], nums)
	return xorSum_with_element + xorSum_without_element
}
