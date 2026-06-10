package main

func allPathsSourceTarget(graph [][]int) [][]int {
	var result [][]int
	target := len(graph) - 1

	var dfs func(node int, path []int)
	dfs = func(node int, path []int) {
		if node == target {
			pathCopy := make([]int, len(path))
			copy(pathCopy, path)
			result = append(result, pathCopy)
			return
		}

		for _, neighbor := range graph[node] {
			path = append(path, neighbor)

			dfs(neighbor, path)

			path = path[:len(path)-1]
		}
	}

	dfs(0, []int{0})

	return result
}
