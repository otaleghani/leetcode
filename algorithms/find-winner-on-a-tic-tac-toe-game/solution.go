package main

func tictactoe(moves [][]int) string {
	rows := make([]int, 3)
	cols := make([]int, 3)

	diag := 0
	antiDiag := 0

	player := 1

	for _, move := range moves {
		r, c := move[0], move[1]

		rows[r] += player
		cols[c] += player

		if r == c {
			diag += player
		}

		if r+c == 2 {
			antiDiag += player
		}

		if abs(rows[r]) == 3 || abs(cols[c]) == 3 || abs(diag) == 3 || abs(antiDiag) == 3 {
			if player == 1 {
				return "A"
			}
			return "B"
		}

		player *= -1
	}

	if len(moves) == 9 {
		return "Draw"
	}

	return "Pending"
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
