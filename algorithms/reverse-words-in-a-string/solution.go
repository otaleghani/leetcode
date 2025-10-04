package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	words := strings.Fields(s)
	result := ""

	for i := len(words) - 1; i >= 0; i-- {
		result += words[i]
		if i != 0 {
			result += " "
		}
	}

	return result
}

// Using strings.Builder uses less memory because you are not creating copies of the string
func reverseWordsAlternative(s string) string {
	words := strings.Fields(s)
	var result strings.Builder

	for i := len(words) - 1; i >= 0; i-- {
		fmt.Fprint(&result, words[i])
		if i != 0 {
			fmt.Fprint(&result, " ")
		}
	}

	return result.String()
}
