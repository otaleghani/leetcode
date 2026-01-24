package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

// This is slightly faster tho
func countNodesAlternative(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countNodes(root.Left)
	right := countNodes(root.Right)

	return 1 + left + right
}
