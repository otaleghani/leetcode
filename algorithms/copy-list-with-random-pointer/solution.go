package main

// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return nil
	}
	nodes := make(map[*Node]*Node)
	headRoot := head

	for head != nil {
		// pair up original with clones
		nodes[head] = &Node{Val: head.Val}
		head = head.Next
	}

	head = headRoot
	copy := nodes[head]
	copyRoot := copy
	for head != nil {
		copy.Next = nodes[head.Next]
		copy.Random = nodes[head.Random]

		head = head.Next
		copy = copy.Next
	}

	return copyRoot
}
