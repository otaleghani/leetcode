package main

func letterCombinations(digits string) []string {
	digitsMap := make(map[rune]string)

	digitsMap['2'] = "abc"
	digitsMap['3'] = "def"
	digitsMap['4'] = "ghi"
	digitsMap['5'] = "jkl"
	digitsMap['6'] = "mno"
	digitsMap['7'] = "pqrs"
	digitsMap['8'] = "tuv"
	digitsMap['9'] = "wxyz"

	var currentDigits []string
	length := 1 // Has the dimension of the result array
	for _, digit := range digits {
		currentDigits = append(currentDigits, digitsMap[digit])
		length = len(digitsMap[digit]) * length
	}
	var result []string

	createCombination(currentDigits, &result, 0, "")

	return result
}

func createCombination(listOfLetters []string, result *[]string, currentHeight int, currentCombination string) {
	if currentHeight > len(listOfLetters)-1 {
		// Add the current combination and delete the last digit
		*result = append(*result, currentCombination)
		// currentCombination = currentCombination[:len(currentCombination)-1]
		return
	}

	letters := listOfLetters[currentHeight]
	for _, letter := range letters {
		currentCombination = currentCombination + string(letter)
		createCombination(listOfLetters, result, currentHeight+1, currentCombination)
		currentCombination = currentCombination[:len(currentCombination)-1]
	}
}
