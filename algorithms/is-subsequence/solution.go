package main

func isSubsequence(s string, t string) bool {
	j := 0

	for i := 0; i < len(t); i++ {
		if j == len(s) {
			return true
		}
		if t[i] == s[j] {
			j++
		}

	}

	if j == len(s) {
		return true
	}
	return false
}
