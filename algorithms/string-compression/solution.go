package main

import "strconv"

func compress(chars []byte) int {
	currentChar := chars[0]
	currentPos := 0
	currentLength := 0

	if len(chars) == 1 {
		return 1
	}

	for i := 0; i < len(chars); i++ {
		if currentChar == chars[i] {
			currentLength++
			continue
		}

		chars[currentPos] = currentChar
		currentPos++
		currentChar = chars[i]
		if currentLength > 1 {
			for _, v := range []byte(strconv.Itoa(currentLength)) {
				chars[currentPos] = v
				currentPos++
			}
			currentLength = 1
		}
	}

	if currentLength > 1 {
		chars[currentPos] = currentChar
		currentPos++
		for _, v := range []byte(strconv.Itoa(currentLength)) {
			chars[currentPos] = v
			currentPos++
		}
	} else {
		chars[currentPos] = chars[len(chars)-1]
		currentPos++
	}

	return currentPos
}

// Using two pointers is more elegant and uses less space
func compressAlternative(chars []byte) int {
	read, write := 0, 0

	for read < len(chars) {
		char := chars[read]
		count := 0

		for read < len(chars) && chars[read] == char {
			read++
			count++
		}

		chars[write] = char
		write++

		if count > 1 {
			for _, v := range []byte(strconv.Itoa(count)) {
				chars[write] = v
				write++
			}
		}
	}

	return write
}
