package main

type Pair struct {
	val int
	pos int
}

func dailyTemperatures(temperatures []int) []int {
	stack := []Pair{}
	res := make([]int, len(temperatures))

	for i := len(temperatures) - 1; i >= 0; i-- {
		if len(stack) == 0 {
			stack = []Pair{Pair{val: temperatures[i], pos: i}}
			res[i] = 0
		} else {
			current := temperatures[i]
			peek := stack[len(stack)-1]
			for len(stack) > 0 && peek.val <= current {
				stack = stack[:len(stack)-1]
				if len(stack) > 0 {
					peek = stack[len(stack)-1]
				}
			}

			if len(stack) == 0 {
				stack = []Pair{Pair{val: temperatures[i], pos: i}}
				res[i] = 0
			} else {
				res[i] = stack[len(stack)-1].pos - i
				stack = append(stack, Pair{val: current, pos: i})
			}
		}

	}

	return res
}
