package main

func shortestBridge(grid [][]int) int {
	n := len(grid)
	queue := [][]int{}
	found := false

	var dfs func(r, c int)
	dfs = func(r, c int) {
		if r < 0 || r >= n || c < 0 || c >= n || grid[r][c] != 1 {
			return
		}
		grid[r][c] = 2
		queue = append(queue, []int{r, c})
		dfs(r+1, c)
		dfs(r-1, c)
		dfs(r, c+1)
		dfs(r, c-1)
	}

	for i := 0; i < n; i++ {
		if found {
			break
		}
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				dfs(i, j)
				found = true
				break
			}
		}
	}

	steps := 0
	dirs := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			for _, d := range dirs {
				nr, nc := curr[0]+d[0], curr[1]+d[1]
				if nr >= 0 && nr < n && nc >= 0 && nc < n {
					if grid[nr][nc] == 1 {
						return steps
					}
					if grid[nr][nc] == 0 {
						grid[nr][nc] = 2
						queue = append(queue, []int{nr, nc})
					}
				}
			}
		}
		steps++
	}

	return -1
}
