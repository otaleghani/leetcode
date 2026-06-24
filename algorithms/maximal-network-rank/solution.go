package main

func maximalNetworkRank(n int, roads [][]int) int {
	degrees := make([]int, n)
	connected := make([][]bool, n)
	for i := range connected {
		connected[i] = make([]bool, n)
	}

	for _, road := range roads {
		u, v := road[0], road[1]
		degrees[u]++
		degrees[v]++
		connected[u][v] = true
		connected[v][u] = true
	}

	maxRank := 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			rank := degrees[i] + degrees[j]
			if connected[i][j] {
				rank--
			}
			if rank > maxRank {
				maxRank = rank
			}
		}
	}

	return maxRank
}
