package main

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) []int {
	graph := make([][][2]int, n)

	for _, edge := range redEdges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], [2]int{v, 0})
	}

	for _, edge := range blueEdges {
		u, v := edge[0], edge[1]
		graph[u] = append(graph[u], [2]int{v, 1})
	}

	ans := make([]int, n)
	for i := range ans {
		ans[i] = -1 // Default value if unreachable
	}
	ans[0] = 0 // Distance to start node is always 0

	q := [][]int{{0, 0, -1}}

	visited := make([][2]bool, n)
	visited[0][0] = true
	visited[0][1] = true

	for len(q) > 0 {
		curr := q[0]
		q = q[1:] // pop

		node, dist, lastColor := curr[0], curr[1], curr[2]

		for _, neighbor := range graph[node] {
			nextNode, nextColor := neighbor[0], neighbor[1]
			if nextColor != lastColor && !visited[nextNode][nextColor] {
				visited[nextNode][nextColor] = true
				if ans[nextNode] == -1 {
					ans[nextNode] = dist + 1
				}

				// Push next state into the queue
				q = append(q, []int{nextNode, dist + 1, nextColor})
			}
		}
	}

	return ans
}
