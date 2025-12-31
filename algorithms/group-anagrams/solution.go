package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	result := make(map[string][]string)

	for i := 0; i < len(strs); i++ {
		chars := []rune(strs[i])
		slices.Sort(chars)
		result[string(chars)] = append(result[string(chars)], strs[i])
	}

	arrResult := [][]string{}
	for _, values := range result {
		if len(values) == 0 {
			continue
		}
		arrResult = append(arrResult, values)
	}

	return arrResult
}
