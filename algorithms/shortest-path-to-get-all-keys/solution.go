package main

func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	startR, startC, numKeys := 0, 0, 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '@' {
				startR, startC = i, j
			} else if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				numKeys++
			}
		}
	}

	target := (1 << numKeys) - 1
	queue := [][]int{{startR, startC, 0, 0}}
	
	visited := make([][][]bool, m)
	for i := range visited {
		visited[i] = make([][]bool, n)
		for j := range visited[i] {
			visited[i][j] = make([]bool, 1<<numKeys)
		}
	}
	visited[startR][startC][0] = true

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		r, c, keys, steps := curr[0], curr[1], curr[2], curr[3]

		if keys == target {
			return steps
		}

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]

			if nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] != '#' {
				cell := grid[nr][nc]
				newKeys := keys

				if cell >= 'A' && cell <= 'F' && (keys&(1<<(cell-'A'))) == 0 {
					continue
				}

				if cell >= 'a' && cell <= 'f' {
					newKeys |= 1 << (cell - 'a')
				}

				if !visited[nr][nc][newKeys] {
					visited[nr][nc][newKeys] = true
					queue = append(queue, []int{nr, nc, newKeys, steps + 1})
				}
			}
		}
	}

	return -1
}
