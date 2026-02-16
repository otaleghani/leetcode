package main

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	graph := make(map[string]map[string]float64)

	for i, eq := range equations {
		u, v := eq[0], eq[1]
		val := values[i]

		if graph[u] == nil {
			graph[u] = make(map[string]float64)
		}
		if graph[v] == nil {
			graph[v] = make(map[string]float64)
		}

		graph[u][v] = val
		graph[v][u] = 1.0 / val
	}

	results := make([]float64, len(queries))

	type state struct {
		node  string  // The current node we're visiting
		value float64 // The cumulative product from start to this node (e.g., start / node)
	}

	for i, query := range queries {
		start, end := query[0], query[1]

		if _, ok := graph[start]; !ok {
			results[i] = -1.0
			continue
		}
		if _, ok := graph[end]; !ok {
			results[i] = -1.0
			continue
		}

		queue := []state{{node: start, value: 1.0}}
		visited := make(map[string]bool)
		visited[start] = true

		found := false

		for len(queue) > 0 {
			curr := queue[0]
			queue = queue[1:]

			if curr.node == end {
				results[i] = curr.value
				found = true
				break
			}

			for neighbor, weight := range graph[curr.node] {
				if !visited[neighbor] {
					visited[neighbor] = true

					newValue := curr.value * weight

					queue = append(queue, state{node: neighbor, value: newValue})
				}
			}
		}

		if !found {
			results[i] = -1.0
		}
	}

	return results
}

func calcEquationAlt(equations [][]string, values []float64, queries [][]string) []float64 {
	nodes := make(map[string]map[string]float64)

	for i, eq := range equations {
		u, v := eq[0], eq[1]
		val := values[i]

		if nodes[u] == nil {
			nodes[u] = make(map[string]float64)
		}
		if nodes[v] == nil {
			nodes[v] = make(map[string]float64)
		}

		nodes[u][v] = val
		nodes[v][u] = 1.0 / val
	}

	var dfs func(curr, target string, visited map[string]bool) float64
	dfs = func(curr, target string, visited map[string]bool) float64 {
		if curr == target {
			return 1.0
		}

		visited[curr] = true

		for neighbor, weight := range nodes[curr] {
			if !visited[neighbor] {
				res := dfs(neighbor, target, visited)
				if res != -1.0 {
					return weight * res
				}
			}
		}

		return -1.0
	}

	results := make([]float64, len(queries))
	for i, q := range queries {
		start, end := q[0], q[1]

		if nodes[start] == nil || nodes[end] == nil {
			results[i] = -1.0
		} else if start == end {
			results[i] = 1.0
		} else {
			visited := make(map[string]bool)
			results[i] = dfs(start, end, visited)
		}
	}

	return results
}
