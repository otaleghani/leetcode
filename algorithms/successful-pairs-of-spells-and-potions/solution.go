package main

import "sort"

func successfulPairs(spells []int, potions []int, success int64) []int {
	result := make([]int, len(spells))
	sort.Ints(potions)
	for idxSpell, spell := range spells {
		start := 0
		end := len(potions) - 1

		for start <= end {
			mid := start + (end-start)/2
			currValue := int64(potions[mid] * spell)
			if currValue >= success {
				result[idxSpell] = len(potions[mid:])
				end = mid - 1
			}
			if currValue < success {
				start = mid + 1
			}
		}
	}
	return result
}
