package main

func rangeAddQueries(n int, queries [][]int) [][]int {
	diffMat := make([][]int, n+1)
	for i := 0; i < n+1; i++ {
		diffMat[i] = make([]int, n+1)
	}

	for _, query := range queries {
		diffMat[query[0]][query[1]] += 1
		diffMat[query[2]+1][query[3]+1] += 1
		diffMat[query[0]][query[3]+1] -= 1
		diffMat[query[2]+1][query[1]] -= 1
	}

	matrix := make([][]int, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			top := 0
			if i != 0 {
				top = matrix[i-1][j]
			}

			left := 0
			if j != 0 {
				left = matrix[i][j-1]
			}

			diagonal := 0
			if i != 0 && j != 0 {
				diagonal = matrix[i-1][j-1]
			}

			matrix[i][j] = diffMat[i][j] + top + left - diagonal
		}
	}

	return matrix
}
