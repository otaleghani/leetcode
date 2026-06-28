package main

func smallestEquivalentString(s1 string, s2 string, baseStr string) string {
	parent := make([]int, 26)
	for i := 0; i < 26; i++ {
		parent[i] = i
	}

	var find func(int) int
	find = func(i int) int {
		if parent[i] == i {
			return i
		}
		parent[i] = find(parent[i])
		return parent[i]
	}

	union := func(i, j int) {
		rootI := find(i)
		rootJ := find(j)
		if rootI < rootJ {
			parent[rootJ] = rootI
		} else if rootI > rootJ {
			parent[rootI] = rootJ
		}
	}

	for i := 0; i < len(s1); i++ {
		union(int(s1[i]-'a'), int(s2[i]-'a'))
	}

	res := make([]byte, len(baseStr))
	for i := 0; i < len(baseStr); i++ {
		res[i] = byte(find(int(baseStr[i]-'a')) + 'a')
	}

	return string(res)
}
