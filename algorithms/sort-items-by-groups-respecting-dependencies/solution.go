package main

func sortItems(n int, m int, group []int, beforeItems [][]int) []int {
	for i := range group {
		if group[i] == -1 {
			group[i] = m
			m++
		}
	}

	itemAdj := make([][]int, n)
	itemInDegree := make([]int, n)
	groupAdj := make([][]int, m)
	groupInDegree := make([]int, m)

	for i := 0; i < n; i++ {
		for _, j := range beforeItems[i] {
			itemAdj[j] = append(itemAdj[j], i)
			itemInDegree[i]++
			if group[i] != group[j] {
				groupAdj[group[j]] = append(groupAdj[group[j]], group[i])
				groupInDegree[group[i]]++
			}
		}
	}

	allItems := make([]int, n)
	for i := 0; i < n; i++ {
		allItems[i] = i
	}

	allGroups := make([]int, m)
	for i := 0; i < m; i++ {
		allGroups[i] = i
	}

	sortedItems := topSort(itemAdj, itemInDegree, allItems)
	if len(sortedItems) == 0 {
		return []int{}
	}

	sortedGroups := topSort(groupAdj, groupInDegree, allGroups)
	if len(sortedGroups) == 0 {
		return []int{}
	}

	groupToItems := make([][]int, m)
	for _, item := range sortedItems {
		g := group[item]
		groupToItems[g] = append(groupToItems[g], item)
	}

	var res []int
	for _, g := range sortedGroups {
		res = append(res, groupToItems[g]...)
	}

	return res
}

func topSort(adj [][]int, inDegree []int, items []int) []int {
	var queue []int
	for _, item := range items {
		if inDegree[item] == 0 {
			queue = append(queue, item)
		}
	}

	var res []int
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		res = append(res, curr)

		for _, next := range adj[curr] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	if len(res) != len(items) {
		return []int{}
	}

	return res
}
