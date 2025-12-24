package main

func isValidSudoku(board [][]byte) bool {
	// Check rows
	for i := range 9 {
		if hasDuplicates(board[i]) {
			return false
		}
	}

	// Check columns
	for i := range 9 {
		col := []byte{}
		for j := range 9 {
			col = append(col, board[j][i])
		}
		if hasDuplicates(col) {
			return false
		}
	}

	// Check square
	for i := range 9 {
		startCol := (i % 3) * 3
		startRow := (i / 3) * 3

		square := []byte{}
		for j := range 3 {
			for k := range 3 {
				square = append(square, board[startRow+j][startCol+k])
			}
		}

		if hasDuplicates(square) {
			return false
		}
	}

	return true
}

func hasDuplicates(slice []byte) bool {
	seen := make(map[byte]bool)

	for _, value := range slice {
		if value == '.' {
			continue
		}
		if seen[value] {
			return true
		}
		seen[value] = true
	}
	return false
}
