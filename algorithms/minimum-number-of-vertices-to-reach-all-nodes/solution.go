package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	inDegree := make([]bool, n)
	for _, edge := range edges {
		inDegree[edge[1]] = true
	}

	var result []int
	for i := 0; i < n; i++ {
		if !inDegree[i] {
			result = append(result, i)
		}
	}

	return result
}
