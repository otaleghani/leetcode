package main

import "strconv"

func evalRPN(tokens []string) int {
	ints := []int{}

	pushInt := func(i int) {
		ints = append(ints, i)
	}
	popInt := func() int {
		if len(ints) == 0 {
			return 0
		}
		last := ints[len(ints)-1]
		ints = ints[:len(ints)-1]
		return last
	}

	for _, token := range tokens {
		switch token {
		case "+", "-", "*", "/":
			second := popInt()
			first := popInt()

			switch token {
			case "+":
				pushInt(second + first)
			case "-":
				pushInt(first - second)
			case "*":
				pushInt(first * second)
			case "/":
				pushInt(first / second)
			}

		default:
			val, _ := strconv.Atoi(token)
			pushInt(val)
		}
	}

	// Ints should have only one at the end
	return popInt()
}
