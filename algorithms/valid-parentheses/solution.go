package main

func isValid(s string) bool {
	stack := []byte{}

	push := func(n byte) {
		stack = append(stack, n)
	}

	pop := func() {
		stack = stack[:len(stack)-1]
	}

	peek := func() byte {
		return stack[len(stack)-1]
	}

	for i := 0; i < len(s); i++ {
		if len(stack) == 0 {
			push(s[i])
			continue
		}

		if s[i] == ')' && peek() == '(' {
			pop()
			continue
		}
		if s[i] == ']' && peek() == '[' {
			pop()
			continue
		}
		if s[i] == '}' && peek() == '{' {
			pop()
			continue
		}

		push(s[i])
	}

	if len(stack) == 0 {
		return true
	}
	return false
}
