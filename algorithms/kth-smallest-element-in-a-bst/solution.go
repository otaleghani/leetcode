package main

import "slices"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	vals := []int{}
	vals = append(vals, root.Val)
	var f func(root *TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			vals = append(vals, root.Left.Val)
		}
		if root.Right != nil {
			vals = append(vals, root.Right.Val)
		}
		f(root.Left)
		f(root.Right)
	}
	f(root)

	slices.Sort(vals)
	return vals[k-1]
}
