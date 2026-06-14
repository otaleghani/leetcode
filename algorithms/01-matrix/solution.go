package main

func updateMatrix(mat [][]int) [][]int {
	m, n := len(mat), len(mat[0])
	queue := [][]int{}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if mat[i][j] == 0 {
				queue = append(queue, []int{i, j})
			} else {
				mat[i][j] = -1
			}
		}
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		r, c := curr[0], curr[1]

		for _, d := range dirs {
			nr, nc := r+d[0], c+d[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n && mat[nr][nc] == -1 {
				mat[nr][nc] = mat[r][c] + 1
				queue = append(queue, []int{nr, nc})
			}
		}
	}

	return mat
}
