package main

func setZeroes(matrix [][]int) {
	cols := make(map[int]struct{})
	rows := make(map[int]struct{})

	// Find the zeroes
	// Change it to another int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				cols[j] = struct{}{}
				rows[i] = struct{}{}
			}
		}
	}

	for col := range cols {
		for i := 0; i < len(matrix); i++ {
			matrix[i][col] = 0
		}
	}

	for row := range rows {
		for j := 0; j < len(matrix[0]); j++ {
			matrix[row][j] = 0
		}
	}
}
