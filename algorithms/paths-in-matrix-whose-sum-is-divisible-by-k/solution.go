package main

import "math"

func numberOfPaths(grid [][]int, k int) int {
	ROWS, COLS := len(grid), len(grid[0])
	MOD := int(math.Pow(float64(10), float64(9))) + 7

	cache := make([][][]int, ROWS)
	for i := range cache {
		cache[i] = make([][]int, COLS)
		for j := range cache[i] {
			cache[i][j] = make([]int, k)
			for z := range cache[i][j] {
				cache[i][j][z] = -1
			}
		}
	}

	var f func(row, col, rem int) int
	f = func(row, col, rem int) int {
		if row == ROWS-1 && col == COLS-1 {
			rem = (rem + grid[row][col]) % k
			if rem != 0 {
				return 0
			} else {
				return 1
			}
		}
		if row == ROWS || col == COLS {
			return 0
		}
		if cache[row][col][rem] > -1 {
			return cache[row][col][rem]
		}

		cache[row][col][rem] = (f(row+1, col, (rem+grid[row][col])%k)%MOD +
			f(row, col+1, (rem+grid[row][col])%k)%MOD) % MOD
		return cache[row][col][rem]
	}

	return f(0, 0, 0) % MOD
}
