func eventualSafeNodes(graph [][]int) []int {
	n := len(graph)
	// colors array: 0 = unvisited, 1 = visiting, 2 = safe
	colors := make([]int, n)

	var dfs func(int) bool
	dfs = func(node int) bool {
		// If the node is already visited, check if it's safe
		if colors[node] != 0 {
			return colors[node] == 2
		}

		// Mark the current node as visiting (part of current path)
		colors[node] = 1

		// Check all neighbors
		for _, neighbor := range graph[node] {
			// If the neighbor is safe, we don't need to visit it again
			if colors[neighbor] == 2 {
				continue
			}
			// If the neighbor is currently being visited (cycle detected)
			// or if a deeper DFS finds a cycle
			if colors[neighbor] == 1 || !dfs(neighbor) {
				return false
			}
		}

		// If all paths lead to safe nodes, mark this node as safe
		colors[node] = 2
		return true
	}

	var result []int
	// Iterate through all nodes to find the safe ones
	for i := 0; i < n; i++ {
		if dfs(i) {
			result = append(result, i)
		}
	}

	// The problem requires the output to be in ascending order.
	// Since we iterate from 0 to n-1, the result is naturally sorted.
	return result
}
