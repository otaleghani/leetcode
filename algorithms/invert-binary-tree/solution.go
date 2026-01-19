package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	tmp := root.Right
	root.Right = root.Left
	root.Left = tmp

	invertTree(root.Right)
	invertTree(root.Left)

	return root
}
