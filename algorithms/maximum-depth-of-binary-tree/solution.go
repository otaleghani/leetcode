package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	max := 1
	if root == nil {
		return max - 1
	}

	leftMax := max + maxDepth(root.Left)
	rightMax := max + maxDepth(root.Right)

	if leftMax > rightMax {
		return leftMax
	}
	return rightMax
}
