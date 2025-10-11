package main

type Stack struct {
	items []byte
}

func (s *Stack) Push(b byte) {
	s.items = append(s.items, b)
}

func (s *Stack) Pop() {
	s.items = s.items[:len(s.items)-1]
}

func removeStars(s string) string {
	stack := Stack{items: []byte{}}
	for _, r := range s {
		if r != '*' {
			stack.Push(byte(r))
			continue
		}
		stack.Pop()
	}

	return string(stack.items)
}
