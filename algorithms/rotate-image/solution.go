package main

func rotate(matrix [][]int) {
	length := len(matrix)

	for i := 0; i < length/2; i++ {
		currMaxLength := length - i - 1

		for j := i; j < currMaxLength; j++ {
			one := matrix[i][j]
			two := matrix[j][currMaxLength]
			three := matrix[currMaxLength][length-j-1]
			four := matrix[length-j-1][i]

			matrix[i][j] = four
			matrix[j][currMaxLength] = one
			matrix[currMaxLength][length-j-1] = two
			matrix[length-j-1][i] = three
		}
	}
}
