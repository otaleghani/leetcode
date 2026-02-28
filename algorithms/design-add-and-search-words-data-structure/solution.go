package main

type WordDictionary struct {
	children  map[rune]*WordDictionary
	exitPoint bool
}

func Constructor() WordDictionary {
	return WordDictionary{children: make(map[rune]*WordDictionary, 26)}
}

func (this *WordDictionary) AddWord(word string) {
	curr := this
	for _, char := range word {
		if _, exists := curr.children[char]; !exists {
			curr.children[char] = &WordDictionary{
				children:  make(map[rune]*WordDictionary, 26),
				exitPoint: false,
			}
		}
		curr = curr.children[char]
	}
	curr.exitPoint = true
}

func (this *WordDictionary) Search(word string) bool {
	queue := []*WordDictionary{this}

	idx := 0
	for len(queue) > 0 {
		size := len(queue)
		var char rune = rune(word[idx])

		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if char == '.' {
				for _, child := range curr.children {
					queue = append(queue, child)
				}
			} else {
				if next, exists := curr.children[char]; exists {
					queue = append(queue, next)
				}
			}
		}
		idx++

		if idx == len(word) {
			for _, val := range queue {
				if val.exitPoint {
					return true
				}
			}
			return false
		}
	}

	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
