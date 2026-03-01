package main

// ==================================================
// Graph solution that I tried initially
// ==================================================

type Node struct {
	x, y      int
	letter    byte
	neighbors []*Node
}

func findWordsGRAPH(board [][]byte, words []string) []string {
	lookup := buildGraph(board)
	res := []string{}

	for _, word := range words {
		// Start search only if the first letter exists on the board
		if startingNodes, exists := lookup[word[0]]; exists {
			for _, startNode := range startingNodes {
				if dfsGraph(startNode, word, 1) { // start at index 1
					res = append(res, word)
					break // Word found, move to the next word
				}
			}
		}
	}
	return res
}

func buildGraph(board [][]byte) map[byte][]*Node {
	lenX := len(board)
	lenY := len(board[0])
	graphBoard := make([][]*Node, lenX)
	lookup := make(map[byte][]*Node)

	// Create all nodes and populate lookup
	for x := 0; x < lenX; x++ {
		graphBoard[x] = make([]*Node, lenY)
		for y := 0; y < lenY; y++ {
			letter := board[x][y]
			node := &Node{x: x, y: y, letter: letter}
			graphBoard[x][y] = node
			lookup[letter] = append(lookup[letter], node)
		}
	}

	// Connect the 4 directional neighbors
	for x := 0; x < lenX; x++ {
		for y := 0; y < lenY; y++ {
			curr := graphBoard[x][y]
			if x > 0 {
				curr.neighbors = append(curr.neighbors, graphBoard[x-1][y])
			} // TOP
			if x < lenX-1 {
				curr.neighbors = append(curr.neighbors, graphBoard[x+1][y])
			} // BOT
			if y > 0 {
				curr.neighbors = append(curr.neighbors, graphBoard[x][y-1])
			} // LEFT
			if y < lenY-1 {
				curr.neighbors = append(curr.neighbors, graphBoard[x][y+1])
			} // RIGHT
		}
	}
	return lookup
}

func dfsGraph(node *Node, word string, index int) bool {
	// Base case: we found the whole word
	if index == len(word) {
		return true
	}

	// Temporarily mark as visited to prevent reusing this cell
	originalLetter := node.letter
	node.letter = '#'

	found := false
	for _, neighbor := range node.neighbors {
		// If neighbor matches the next letter and hasn't been visited
		if neighbor.letter == word[index] {
			if dfsGraph(neighbor, word, index+1) {
				found = true
				break
			}
		}
	}

	// Backtrack to restore the letter
	node.letter = originalLetter
	return found
}

// ==================================================
// Standard trie solution
// ==================================================

type TrieNode struct {
	children [26]*TrieNode
	word     string // Store the full word at the end node
}

func findWords(board [][]byte, words []string) []string {
	// Build the Trie
	root := &TrieNode{}
	for _, word := range words {
		curr := root
		for i := 0; i < len(word); i++ {
			idx := word[i] - 'a'
			if curr.children[idx] == nil {
				curr.children[idx] = &TrieNode{}
			}
			curr = curr.children[idx]
		}
		curr.word = word // Mark the end of a word
	}

	res := []string{}
	lenX, lenY := len(board), len(board[0])

	// Run DFS from every cell on the board
	var dfs func(x, y int, node *TrieNode)
	dfs = func(x, y int, node *TrieNode) {
		// Bounds check and visited check
		if x < 0 || x >= lenX || y < 0 || y >= lenY || board[x][y] == '#' {
			return
		}

		char := board[x][y]
		nextNode := node.children[char-'a']

		// If this path isn't in our Trie, stop searching immediately (Pruning)
		if nextNode == nil {
			return
		}

		// If we found a word, add it and clear it to prevent duplicates
		if nextNode.word != "" {
			res = append(res, nextNode.word)
			nextNode.word = ""
		}

		// Mark visited
		board[x][y] = '#'

		// Explore 4 directions
		dfs(x-1, y, nextNode) // TOP
		dfs(x+1, y, nextNode) // BOT
		dfs(x, y-1, nextNode) // LEFT
		dfs(x, y+1, nextNode) // RIGHT

		// Backtrack
		board[x][y] = char
	}

	for i := 0; i < lenX; i++ {
		for j := 0; j < lenY; j++ {
			dfs(i, j, root)
		}
	}

	return res
}
