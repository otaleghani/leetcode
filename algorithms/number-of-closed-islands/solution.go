package main

func closedIsland(grid [][]int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    m, n := len(grid), len(grid[0])

    var dfs func(i, j int)
    dfs = func(i, j int) {
        if i < 0 || i >= m || j < 0 || j >= n || grid[i][j] != 0 {
            return
        }
        grid[i][j] = 1
        dfs(i+1, j)
        dfs(i-1, j)
        dfs(i, j+1)
        dfs(i, j-1)
    }

    for i := 0; i < m; i++ {
        if grid[i][0] == 0 {
            dfs(i, 0)
        }
        if grid[i][n-1] == 0 {
            dfs(i, n-1)
        }
    }

    for j := 0; j < n; j++ {
        if grid[0][j] == 0 {
            dfs(0, j)
        }
        if grid[m-1][j] == 0 {
            dfs(m-1, j)
        }
    }

    count := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 0 {
                count++
                dfs(i, j)
            }
        }
    }

    return count
}
