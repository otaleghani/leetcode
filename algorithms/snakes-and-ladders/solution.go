package main

func snakesAndLadders(board [][]int) int {
	n := len(board)
	target := n * n

	getCoordinates := func(square int) (int, int) {
		zeroIndexed := square - 1
		rowFromBottom := zeroIndexed / n
		row := (n - 1) - rowFromBottom

		col := zeroIndexed % n
		// If the row from the bottom is odd, the board direction reverses (right to left)
		if rowFromBottom%2 != 0 {
			col = (n - 1) - col
		}
		return row, col
	}

	queue := []int{1} // Start at square 1
	visited := make([]bool, target+1)
	visited[1] = true
	moves := 0

	for len(queue) > 0 {
		size := len(queue)

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			// If we reached the end, return the number of dice rolls it took
			if curr == target {
				return moves
			}

			// Roll the die: 1 through 6
			for roll := 1; roll <= 6; roll++ {
				nextSquare := curr + roll
				if nextSquare > target {
					break // Can't move past the final square
				}

				// Peek at the board to see if there is a snake/ladder
				r, c := getCoordinates(nextSquare)
				destination := nextSquare
				if board[r][c] != -1 {
					// The number on the board is the destination
					destination = board[r][c]
				}

				// If we haven't visited this destination yet, add it to the queue
				if !visited[destination] {
					visited[destination] = true
					queue = append(queue, destination)
				}
			}
		}
		moves++
	}

	return -1
}
