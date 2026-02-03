package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	levels := [][]int{}

	var f func(*TreeNode, int)
	f = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(levels)-1 < level {
			levels = append(levels, []int{})
		}
		levels[level] = append(levels[level], root.Val)

		f(root.Left, level+1)
		f(root.Right, level+1)
	}
	f(root, 0)

	return levels
}
