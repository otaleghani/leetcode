package main

// Definition for a Node.
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return nil
	}

	depthMap := make(map[int][]*Node)

	var addToMap func(root *Node, depth int)
	addToMap = func(root *Node, depth int) {
		if root == nil {
			return
		}

		// Append current node to the slice for this depth
		depthMap[depth] = append(depthMap[depth], root)

		addToMap(root.Left, depth+1)
		addToMap(root.Right, depth+1)
	}

	addToMap(root, 0)

	// Iterate through all depths found
	for _, nodes := range depthMap {
		for i := 0; i < len(nodes)-1; i++ {
			nodes[i].Next = nodes[i+1]
		}
	}

	return root
}
