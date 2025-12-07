package main

import "strings"

func lengthOfLastWord(s string) int {
	trimmed := strings.Trim(s, " ")
	arr := strings.Split(trimmed, " ")
	return len(arr[len(arr)-1])
}
