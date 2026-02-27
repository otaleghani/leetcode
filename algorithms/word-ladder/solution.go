package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	queue := []string{beginWord}
	visited := make(map[string]struct{})
	level := 1
	hasEndWord := false

	for i := 0; i < len(wordList); i++ {
		if wordList[i] == endWord {
			hasEndWord = true
			break
		}
	}

	if !hasEndWord {
		return 0
	}

	visited[beginWord] = struct{}{}

	for len(queue) > 0 {
		size := len(queue)

		for k := 0; k < size; k++ {
			curr := queue[0]
			queue = queue[1:]

			if curr == endWord {
				return level
			}

			for i := 0; i < len(wordList); i++ {
				if _, ok := visited[wordList[i]]; ok {
					continue
				}
				if diff(curr, wordList[i]) == 1 {
					queue = append(queue, wordList[i])
					visited[wordList[i]] = struct{}{}
				}
			}
		}

		level++
	}

	return 0
}

func diff(w1, w2 string) int {
	diff := 0
	for i := 0; i < len(w1); i++ {
		if w1[i] != w2[i] {
			diff++
		}
	}
	return diff
}
