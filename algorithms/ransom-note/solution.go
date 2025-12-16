package main

func canConstruct(ransomNote string, magazine string) bool {
	letters := make([]int, 26)

	for _, letter := range ransomNote {
		letters[letter-97] -= 1
	}

	for _, letter := range magazine {
		letters[letter-97] += 1
	}

	for _, letter := range letters {
		if letter < 0 {
			return false
		}
	}

	return true
}
