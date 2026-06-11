package main

func criticalConnections(n int, connections [][]int) [][]int {
	graph := make([][]int, n)
	for _, conn := range connections {
		u, v := conn[0], conn[1]
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}

	disc := make([]int, n)
	low := make([]int, n)
	for i := 0; i < n; i++ {
		disc[i] = -1
		low[i] = -1
	}

	var res [][]int
	timer := 0

	var dfs func(curr, prev int)
	dfs = func(curr, prev int) {
		disc[curr] = timer
		low[curr] = timer
		timer++

		for _, next := range graph[curr] {
			if next == prev {
				continue
			}
			if disc[next] != -1 {
				if disc[next] < low[curr] {
					low[curr] = disc[next]
				}
			} else {
				dfs(next, curr)
				if low[next] < low[curr] {
					low[curr] = low[next]
				}
				if low[next] > disc[curr] {
					res = append(res, []int{curr, next})
				}
			}
		}
	}

	dfs(0, -1)

	return res
}
