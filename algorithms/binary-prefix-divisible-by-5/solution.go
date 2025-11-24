package main

func prefixesDivBy5(nums []int) []bool {
	result := make([]bool, len(nums))
	remainder := 0

	for i, bit := range nums {
		remainder = (remainder << 1) | bit
		remainder = remainder % 5

		if remainder == 0 {
			result[i] = true
		}
	}

	return result
}
