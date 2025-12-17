package main

import "strings"

func fullJustify(words []string, maxWidth int) []string {
	currLine := []string{}
	currLen := 0
	result := []string{}

	for _, word := range words {
		if currLen+len(word) <= maxWidth {
			currLine = append(currLine, word)
			currLen += len(word) + 1
			continue
		} else {
			spacesToAdd := maxWidth - (currLen - len(currLine))

			// Single word, add the padding at the end
			if len(currLine) == 1 {
				for range spacesToAdd {
					currLine[0] += " "
				}
			} else {
				i := 0
				for spacesToAdd > 0 {
					currIdx := i % len(currLine)
					if currIdx == len(currLine)-1 {
						i++
						continue
					}
					currLine[currIdx] += " "
					spacesToAdd--
					i++
				}
			}

			line := strings.Join(currLine, "")
			result = append(result, line)

			// Then add the current word to an emptied out currLine
			currLine = []string{word}
			currLen = len(word) + 1
		}
	}

	if len(currLine) != 0 {
		spacesToAdd := maxWidth - (currLen - len(currLine)) - len(currLine) + 1
		for range spacesToAdd {
			currLine[len(currLine)-1] += " "
		}
		line := strings.Join(currLine, " ")
		result = append(result, line)
	}

	return result
}
