package main

import "errors"

type Stack struct {
	array []int
}

func (s *Stack) Peek() (int, error) {
	if len(s.array) == 0 {
		return 0, errors.New("No asteroids in array")
	}
	return s.array[len(s.array)-1], nil
}

func (s *Stack) Pop() (int, error) {
	if len(s.array) == 0 {
		return 0, errors.New("No asteroids in array")
	}
	val := s.array[len(s.array)-1]
	s.array = s.array[:len(s.array)-1]
	return val, nil
}

func (s *Stack) Push(v int) {
	s.array = append(s.array, v)
}

func isRight(i int) bool {
	if i > 0 {
		return true
	}
	return false
}

func Abs(i int) int {
	if i < 0 {
		return i * -1
	}
	return i
}

func asteroidCollision(asteroids []int) []int {
	s := Stack{}
	for _, enter := range asteroids {
		exit, err := s.Peek()
		if err != nil {
			s.Push(enter)
			continue
		}

		// Not same direction
		if (isRight(exit) && isRight(enter)) || (!isRight(exit) && !isRight(enter)) {
			s.Push(enter)
			continue
		}

		// Not opposite direction
		if isRight(enter) && !isRight(exit) {
			s.Push(enter)
			continue
		}

		// The will collide!
		// They are the same, delete both
		if Abs(enter) == Abs(exit) {
			s.Pop()
			continue
		}

		for Abs(exit) < Abs(enter) {
			_, err = s.Pop()
			if err != nil {
				s.Push(enter)
				break
			}
			exit, err = s.Peek()
			if isRight(exit) && !isRight(enter) && Abs(enter) == Abs(exit) {
				s.Pop()
				break
			}
			if !(isRight(exit) && !isRight(enter)) {
				s.Push(enter)
				break
			}
		}
	}

	return s.array
}
