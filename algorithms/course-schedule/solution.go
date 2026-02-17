package main

func canFinish(numCourses int, prerequisites [][]int) bool {
	adj := make([][]int, numCourses)
	inDegree := make([]int, numCourses)

	for _, pr := range prerequisites {
		course, prep := pr[0], pr[1]
		adj[prep] = append(adj[prep], course)
		inDegree[course]++
	}

	queue := []int{}
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}

	count := 0
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		count++

		for _, nextCourse := range adj[current] {
			inDegree[nextCourse]--

			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}

	return count == numCourses
}
