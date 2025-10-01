package main

func maximumStrongPairXor(nums []int) int {
	result := 0
	for left := 0; left < len(nums); left++ {
		for right := left; right < len(nums); right++ {
			abs := Abs(nums[left] - nums[right])
			if abs <= min(nums[left], nums[right]) {
				if (nums[left] ^ nums[right]) > result {
					result = nums[left] ^ nums[right]
				}
			}
		}
	}
	return result
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
