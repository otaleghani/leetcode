package main

import "math"

func minimumRecolors(blocks string, k int) int {
	countB := 0
	countW := 0
	left := 0
	result := math.MaxInt

	for right := 0; right < len(blocks); right++ {
		if blocks[right] == 'B' {
			countB++
		} else {
			countW++
		}

		if right-left == k-1 {
			if countW < result {
				result = countW
			}

			// Move one to the left, remove one
			if blocks[left] == 'B' {
				countB--
			} else {
				countW--
			}
			left++
		}
	}

	return result
}
