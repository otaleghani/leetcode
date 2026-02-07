package main

func numIslands(grid [][]byte) int {
	numIslands := 0

	var dfs func(x, y int)
	dfs = func(x, y int) {
		if x > len(grid)-1 || y > len(grid[0])-1 || x < 0 || y < 0 || grid[x][y] != '1' {
			return
		}

		grid[x][y] = '2'

		dfs(x+1, y)
		dfs(x-1, y)
		dfs(x, y+1)
		dfs(x, y-1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				numIslands += 1
				dfs(i, j)
			}
		}
	}

	return numIslands
}
