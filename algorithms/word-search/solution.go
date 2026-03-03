package main

func exist(board [][]byte, word string) bool {
	rows := len(board)
	cols := len(board[0])

	var dfs func(r, c, i int) bool
	dfs = func(r, c, i int) bool {
		if i == len(word) {
			return true
		}

		// Out of bounds or character mismatch
		if r < 0 || r >= rows || c < 0 || c >= cols || board[r][c] != word[i] {
			return false
		}

		// Temporarily mark the cell as visited using a non-alphabet character
		temp := board[r][c]
		board[r][c] = '#'

		// Explore all 4 directions (Up, Down, Right, Left)
		res := dfs(r+1, c, i+1) ||
			dfs(r-1, c, i+1) ||
			dfs(r, c+1, i+1) ||
			dfs(r, c-1, i+1)

		// Restore the cell's original value
		board[r][c] = temp

		return res
	}

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if board[r][c] == word[0] && dfs(r, c, 0) {
				return true
			}
		}
	}

	return false
}
