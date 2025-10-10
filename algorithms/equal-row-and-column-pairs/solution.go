package main

import "strconv"

func equalPairs(grid [][]int) int {
	cols := make(map[string]int)
	rows := make(map[string]int)
	result := 0

	for i := 0; i < len(grid); i++ {
		colIdx := ""
		rowIdx := ""
		for j := 0; j < len(grid[i]); j++ {
			rowIdx += strconv.Itoa(grid[i][j]) + "-"
			colIdx += strconv.Itoa(grid[j][i]) + "-"
		}
		cols[colIdx]++
		rows[rowIdx]++
	}

	for colIdx, colVal := range cols {
		if rowVal, ok := rows[colIdx]; ok {
			result += colVal * rowVal
		}
	}

	return result
}

// You could use a trie for this tho
// Or you could use a polynomial hash to make the hashing better
