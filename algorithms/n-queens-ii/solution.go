package main

func totalNQueens(n int) int {
	cols := make([]bool, n)
	diag1 := make([]bool, 2*n-1) // Anti-diagonals (row + col)
	diag2 := make([]bool, 2*n-1) // Main diagonals (row - col + n - 1)

	count := 0

	var backtrack func(row int)
	backtrack = func(row int) {
		// If we successfully reached the last row, we found a solution
		if row == n {
			count++
			return
		}

		// Try placing a queen in each column of the current row
		for col := 0; col < n; col++ {
			d1 := row + col
			d2 := row - col + n - 1

			// Check if the current position is under attack
			if cols[col] || diag1[d1] || diag2[d2] {
				continue
			}

			// Place the queen (mark the column and diagonals as occupied)
			cols[col] = true
			diag1[d1] = true
			diag2[d2] = true

			// Move to the next row
			backtrack(row + 1)

			// Remove the queen (backtrack to explore other configurations)
			cols[col] = false
			diag1[d1] = false
			diag2[d2] = false
		}
	}

	backtrack(0)

	return count
}
