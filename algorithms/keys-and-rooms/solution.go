package main

func canVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	seen[0] = true
	stack := Stack{array: []int{0}}

	for len(stack.array) != 0 {
		key := stack.Pop()
		for _, val := range rooms[key] {
			if !seen[val] {
				seen[val] = true
				stack.Push(val)
			}
		}
	}

	for _, val := range seen {
		if !val {
			return false
		}
	}
	return true
}

type Stack struct {
	array []int
}

func (s *Stack) Pop() int {
	if len(s.array) == 0 {
		return 0
	}
	val := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return val
}

func (s *Stack) Push(v int) {
	s.array = append(s.array, v)
}
