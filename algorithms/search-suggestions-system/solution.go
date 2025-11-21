package main

import "slices"

func suggestedProducts(products []string, searchWord string) [][]string {
	trie := Constructor()
	//slices.Sort(products)

	for _, product := range products {
		trie.Insert(product)
	}

	result := [][]string{}
	currentSearch := ""
	for _, char := range searchWord {
		currentSearch += string(char)
		result = append(result, trie.Extract(currentSearch))
	}

	return result
}

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

func (this *Trie) Extract(prefix string) []string {
	currentNode := this
	result := []string{}
	for _, char := range prefix {
		if _, exists := currentNode.children[char]; !exists {
			return result
		}
		currentNode = currentNode.children[char]
	}

	// Current node is the prefix
	currentNode.getExit(prefix, &result)
	slices.Sort(result)
	if len(result) > 3 {
		return result[:3]
	}

	return result
}

func (this *Trie) getExit(currentString string, data *[]string) {
	if this.exitPoint {
		*data = append(*data, currentString)
	}
	for char, child := range this.children {
		child.getExit(currentString+string(char), data)
	}
}
