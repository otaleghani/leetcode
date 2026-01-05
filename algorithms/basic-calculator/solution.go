package main

func calculate(s string) int {
	stack := []byte{}
	stackVals := []int{}

	push := func(n byte) {
		stack = append(stack, n)
	}
	pop := func() byte {
		last := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		return last
	}

	pushVals := func(n int) {
		stackVals = append(stackVals, n)
	}
	popVals := func() int {
		last := stackVals[len(stackVals)-1]
		stackVals = stackVals[:len(stackVals)-1]
		return last
	}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case ' ':
			continue

		case '(':
			push(s[i])

		case ')':
			// Pop until we find '('
			var val byte
			values := []int{}
			for {
				val = pop()
				if val == '(' {
					break
				}

				intVal := 0
				if val == 'n' {
					intVal = popVals()
				}

				values = append(values, intVal)
			}

			evalVal := 0
			for _, v := range values {
				evalVal += v
			}

			// Check the sign
			if len(stack) > 0 {
				sign := pop()
				if sign == '-' {
					evalVal *= -1
				} else {
					push(sign)
				}
			}

			push('n') // Flag to get the value from the other stack
			pushVals(evalVal)

		case '+', '-':
			push(s[i])

		default:
			// Parse it and push it. First off find all the digits
			digits := []byte{s[i]}
			for {
				if i+1 < len(s) && isDigit(s[i+1]) {
					digits = append(digits, s[i+1])
					i++
				} else {
					break
				}
			}

			// Parse digits into an int
			digit := 0
			for j := 0; j < len(digits); j++ {
				digit = digit*10 + int(digits[j]-'0')
			}

			var sign byte
			if len(stack) > 0 {
				sign = pop()
			}
			if sign == '-' {
				digit *= -1
				//push('+')
			} else {
				push(sign)
			}

			push('n')
			pushVals(digit)
		}
	}

	result := 0
	for _, v := range stackVals {
		result += v
	}

	return result
}

func isDigit(b byte) bool {
	switch b {
	case ' ', '-', '+', '(', ')':
		return false
	default:
		return true
	}
}
