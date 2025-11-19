package main

func findFinalValue(nums []int, original int) int {
	myMap := make(map[int]struct{}, len(nums))
	for i := range len(nums) {
		myMap[nums[i]] = struct{}{}
	}

	for {
		_, isPresent := myMap[original]
		if isPresent {
			original *= 2
		} else {
			break
		}
	}

	return original
}
