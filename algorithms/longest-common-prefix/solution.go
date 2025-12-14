package main

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	result := ""
	for i := 0; i < len(strs[0]); i++ {
		currChar := strs[0][i]
		for _, str := range strs {
			if i >= len(str) {
				return result
			}
			if str[i] == currChar {
				continue
			} else {
				return result
			}
		}
		result += string(currChar)
	}

	return result
}
