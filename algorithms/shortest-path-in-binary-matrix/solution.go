package main

func shortestPathBinaryMatrix(grid [][]int) int {
	n := len(grid)
	if grid[0][0] == 1 || grid[n-1][n-1] == 1 {
		return -1
	}
	if n == 1 {
		return 1
	}

	queue := [][]int{{0, 0, 1}}
	grid[0][0] = 1

	dirs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		r, c, d := curr[0], curr[1], curr[2]

		for _, dir := range dirs {
			nr, nc := r+dir[0], c+dir[1]

			if nr >= 0 && nr < n && nc >= 0 && nc < n && grid[nr][nc] == 0 {
				if nr == n-1 && nc == n-1 {
					return d + 1
				}
				grid[nr][nc] = 1
				queue = append(queue, []int{nr, nc, d + 1})
			}
		}
	}

	return -1
}
