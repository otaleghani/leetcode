package main

func gameOfLife(board [][]int) {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			countAlive := 0
			rowBefore := i - 1
			rowAfter := i + 1
			colBefore := j - 1
			colAfter := j + 1

			// TOP
			if rowBefore >= 0 {
				// TOP-LEFT
				if colBefore >= 0 && board[rowBefore][colBefore] >= 1 {
					countAlive += 1
				}
				// TOP-CENTER
				if board[rowBefore][j] >= 1 {
					countAlive += 1
				}
				// TOP-RIGHT
				if colAfter <= len(board[i])-1 && board[rowBefore][colAfter] >= 1 {
					countAlive += 1
				}
			}

			// CENTER-LEFT
			if colBefore >= 0 && board[i][colBefore] >= 1 {
				countAlive += 1
			}
			// CENTER-RIGHT
			if colAfter <= len(board[i])-1 && board[i][colAfter] >= 1 {
				countAlive += 1
			}

			// BOT
			if rowAfter <= len(board)-1 {
				// BOT-LEFT
				if colBefore >= 0 && board[rowAfter][colBefore] >= 1 {
					countAlive += 1
				}
				// BOT-CENTER
				if board[rowAfter][j] >= 1 {
					countAlive += 1
				}
				// BOT-RIGHT
				if colAfter <= len(board[i])-1 && board[rowAfter][colAfter] >= 1 {
					countAlive += 1
				}
			}

			if board[i][j] == 1 {
				if countAlive < 2 {
					board[i][j] = 2
				}
				if countAlive == 2 || countAlive == 3 {
					board[i][j] = 1
				}
				if countAlive > 3 {
					board[i][j] = 2
				}
			} else {
				if countAlive == 3 {
					board[i][j] = -1
				}
			}
		}
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == 2 {
				board[i][j] = 0
			}
			if board[i][j] == -1 {
				board[i][j] = 1
			}
		}
	}
}
