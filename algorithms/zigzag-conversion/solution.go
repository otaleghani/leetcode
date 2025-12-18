package main

func convert(s string, numRows int) string {
	if numRows == 1 || numRows >= len(s) {
		return s
	}

	rows := make([][]byte, numRows)

	currRow := 0
	goingDown := false

	for i := 0; i < len(s); i++ {
		rows[currRow] = append(rows[currRow], s[i])

		// If we hit the top or bottom row, reverse the direction
		if currRow == 0 || currRow == numRows-1 {
			goingDown = !goingDown
		}

		if goingDown {
			currRow++
		} else {
			currRow--
		}
	}

	result := make([]byte, 0, len(s))
	for _, row := range rows {
		result = append(result, row...)
	}

	return string(result)
}
