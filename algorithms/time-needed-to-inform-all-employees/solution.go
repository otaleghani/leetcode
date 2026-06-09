package main

func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
	adjList := make([][]int, n)
	for i, mgr := range manager {
		if mgr != -1 {
			adjList[mgr] = append(adjList[mgr], i)
		}
	}

	var dfs func(node int) int
	dfs = func(node int) int {
		maxTime := 0

		// Traverse all direct subordinates
		for _, sub := range adjList[node] {
			time := dfs(sub)
			if time > maxTime {
				maxTime = time
			}
		}

		return informTime[node] + maxTime
	}

	return dfs(headID)
}
