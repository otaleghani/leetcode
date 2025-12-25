package main

func spiralOrder(matrix [][]int) []int {
	result := []int{}

	// Loop until the matrix is empty
	for len(matrix) > 0 {

		// TOP
		result = append(result, matrix[0]...)
		matrix = matrix[1:]
		if len(matrix) == 0 {
			break
		}

		// RIGHT (Top to Bottom)
		for i := 0; i < len(matrix); i++ {
			if len(matrix[i]) == 0 {
				continue
			}
			lastIdx := len(matrix[i]) - 1
			result = append(result, matrix[i][lastIdx])
			matrix[i] = matrix[i][:lastIdx]
		}
		if len(matrix) == 0 || len(matrix[0]) == 0 {
			break
		}

		// BOTTOM
		lastRowIdx := len(matrix) - 1
		lastRow := matrix[lastRowIdx]
		for i := len(lastRow) - 1; i >= 0; i-- {
			result = append(result, lastRow[i])
		}
		matrix = matrix[:lastRowIdx]

		if len(matrix) == 0 {
			break
		}

		// LEFT (Bottom to Top)
		for i := len(matrix) - 1; i >= 0; i-- {
			if len(matrix[i]) == 0 {
				continue
			}
			result = append(result, matrix[i][0])

			matrix[i] = matrix[i][1:]
		}
	}

	return result
}
