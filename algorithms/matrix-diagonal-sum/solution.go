package main

func diagonalSum(mat [][]int) int {
	sum := 0
	for i := 0; i < len(mat); i++ {
		sum += mat[i][i] + mat[i][len(mat[i])-1-i]
		if len(mat[i])%2 != 0 && len(mat[i])/2 == i {
			sum -= mat[i][i]
		}
	}
	return sum
}
