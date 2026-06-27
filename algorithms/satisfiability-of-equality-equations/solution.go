package main

func equationsPossible(equations []string) bool {
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

	for _, eq := range equations {
		if eq[1] == '=' {
			x := int(eq[0] - 'a')
			y := int(eq[3] - 'a')
			parent[find(x)] = find(y)
		}
	}

	for _, eq := range equations {
		if eq[1] == '!' {
			x := int(eq[0] - 'a')
			y := int(eq[3] - 'a')
			if find(x) == find(y) {
				return false
			}
		}
	}

	return true
}
