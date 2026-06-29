package main

func numSimilarGroups(strs []string) int {
	n := len(strs)
	parent := make([]int, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}

	groups := n

	var find func(i int) int
	find = func(i int) int {
		if parent[i] != i {
			parent[i] = find(parent[i])
		}
		return parent[i]
	}

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if isSimilar(strs[i], strs[j]) {
				p1, p2 := find(i), find(j)
				if p1 != p2 {
					parent[p1] = p2
					groups--
				}
			}
		}
	}

	return groups
}

func isSimilar(s1, s2 string) bool {
	diff := 0
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			diff++
			if diff > 2 {
				return false
			}
		}
	}
	return true
}
