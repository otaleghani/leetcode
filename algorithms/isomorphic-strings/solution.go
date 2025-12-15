package main

func isIsomorphic(s string, t string) bool {
	iso := make(map[byte]byte)
	added := make(map[byte]struct{})

	for i := range s {
		val, ok := iso[s[i]]
		if !ok {
			iso[s[i]] = t[i]
			if _, ok := added[t[i]]; ok {
				return false
			}
			added[t[i]] = struct{}{}
			continue
		}
		if val != t[i] {
			return false
		}
	}

	return true
}
