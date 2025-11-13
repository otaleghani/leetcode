package main

func maxOperations(s string) int {
	totalOps := 0
	onesFound := 0
	last := rune(s[0])

	for _, char := range s {
		if char == '1' {
			onesFound += 1
		} else if char == '0' && char != last {
			totalOps += onesFound
		}
		last = char
	}

	return totalOps
}
