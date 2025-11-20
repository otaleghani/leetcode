package main

type Trie struct {
	children  map[rune]*Trie
	exitPoint bool
}

func Constructor() Trie {
	return Trie{children: make(map[rune]*Trie, 26)}
}

func (this *Trie) Insert(word string) {
	currentNode := this
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			// Create the new node
			currentNode.children[char] = &Trie{
				children:  make(map[rune]*Trie, 26),
				exitPoint: false,
			}
		}
		currentNode = currentNode.children[char]
	}
	currentNode.exitPoint = true
}

func (this *Trie) Search(word string) bool {
	currentNode := this
	for _, char := range word {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	if currentNode.exitPoint {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	currentNode := this
	for _, char := range prefix {
		if _, exists := currentNode.children[char]; !exists {
			return false
		}
		currentNode = currentNode.children[char]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
