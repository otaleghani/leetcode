package main

func mergeAlternately(word1 string, word2 string) string {
	result := ""
	length := 0
	remainder := ""

	if len(word1) > len(word2) {
		length = len(word2)
		remainder = word1[length:]
	} else {
		length = len(word1)
		remainder = word2[length:]
	}

	for i := 0; i < length; i++ {
		result += string(word1[i])
		result += string(word2[i])
	}

	return result + remainder
}
