package main

import (
	"strconv"
	"unicode"
)

type Stack struct {
	arr []byte
}

func (s *Stack) Push(b byte) {
	s.arr = append(s.arr, b)
}

func (s *Stack) Pop() byte {
	data := s.arr[len(s.arr)-1]
	s.arr = s.arr[:len(s.arr)-1]
	return data
}

func decodeString(s string) string {
	stack := Stack{}

	for _, v := range s {
		if v != ']' {
			stack.Push(byte(v))
		} else {
			substring := ""

			for len(stack.arr) > 0 && stack.arr[len(stack.arr)-1] != '[' {
				substring = string(stack.Pop()) + substring
			}
			stack.Pop()

			k := ""

			for len(stack.arr) > 0 && unicode.IsDigit(rune(stack.arr[len(stack.arr)-1])) {
				k = string(stack.Pop()) + k
			}

			count := 1
			if k != "" {
				count, _ = strconv.Atoi(k)
			}

			for i := 0; i < count; i++ {
				for _, char := range substring {
					stack.Push(byte(char))
				}
			}
		}
	}

	return string(stack.arr)
}
