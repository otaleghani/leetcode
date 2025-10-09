package main

func findDifference(nums1 []int, nums2 []int) [][]int {
	numsMap1 := make(map[int]struct{})
	numsMap2 := make(map[int]struct{})

	for i := 0; i < len(nums1); i++ {
		numsMap1[nums1[i]] = struct{}{}
	}

	for i := 0; i < len(nums2); i++ {
		numsMap2[nums2[i]] = struct{}{}
	}

	answer1 := []int{}
	answer2 := []int{}
	for i, _ := range numsMap1 {
		if _, ok := numsMap2[i]; !ok {
			answer1 = append(answer1, i)
		}
	}
	for i, _ := range numsMap2 {
		if _, ok := numsMap1[i]; !ok {
			answer2 = append(answer2, i)
		}
	}

	answer := [][]int{answer1, answer2}
	return answer
}

// You can specify the len of maps and the result array, but that's about it
func findDifferenceAlternative(nums1 []int, nums2 []int) [][]int {
	numsMap1 := make(map[int]struct{}, len(nums1))
	numsMap2 := make(map[int]struct{}, len(nums2))

	for i := 0; i < len(nums1); i++ {
		numsMap1[nums1[i]] = struct{}{}
	}

	for i := 0; i < len(nums2); i++ {
		numsMap2[nums2[i]] = struct{}{}
	}

	answer := make([][]int, 2, 2)
	for i, _ := range numsMap1 {
		if _, ok := numsMap2[i]; !ok {
			answer[0] = append(answer[0], i)
		}
	}
	for i, _ := range numsMap2 {
		if _, ok := numsMap1[i]; !ok {
			answer[1] = append(answer[1], i)
		}
	}

	return answer
}
