package main

import "sort"

func accountsMerge(accounts [][]string) [][]string {
	parent := make(map[string]string)
	emailToName := make(map[string]string)

	var find func(string) string
	find = func(s string) string {
		if parent[s] != s {
			parent[s] = find(parent[s])
		}
		return parent[s]
	}

	union := func(s1, s2 string) {
		root1 := find(s1)
		root2 := find(s2)
		if root1 != root2 {
			parent[root1] = root2
		}
	}

	for _, account := range accounts {
		name := account[0]
		for i := 1; i < len(account); i++ {
			email := account[i]
			if _, exists := parent[email]; !exists {
				parent[email] = email
			}
			emailToName[email] = name
			union(account[1], email)
		}
	}

	merged := make(map[string][]string)
	for email := range parent {
		root := find(email)
		merged[root] = append(merged[root], email)
	}

	var res [][]string
	for _, emails := range merged {
		sort.Strings(emails)
		account := append([]string{emailToName[emails[0]]}, emails...)
		res = append(res, account)
	}

	return res
}
