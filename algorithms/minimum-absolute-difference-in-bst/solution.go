package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getMinimumDifference(root *TreeNode) int {
	min := 1000000000
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			diff := root.Val - root.Left.Val
			if diff < min {
				min = diff
			}
			if root.Left.Right != nil {
				r := root.Left
				for r.Right != nil {
					r = r.Right
				}
				diff = root.Val - r.Val
				if diff < min {
					min = diff
				}
			}
		}
		if root.Right != nil {
			diff := root.Right.Val - root.Val
			if diff < min {
				min = diff
			}
			if root.Right.Left != nil {
				r := root.Right
				for r.Left != nil {
					r = r.Left
				}
				diff = r.Val - root.Val
				if diff < min {
					min = diff
				}
			}
		}

		f(root.Left)
		f(root.Right)
	}
	f(root)
	return min
}
