package main

func strStr(haystack string, needle string) int {
	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle[0] {
			found := true
			for j := 1; j < len(needle); j++ {
				if i+j >= len(haystack) || haystack[i+j] != needle[j] {
					found = false
					break
				}
			}
			if found {
				return i
			}
		}
	}
	return -1
}
