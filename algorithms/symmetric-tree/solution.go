package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var check func(left, right *TreeNode) bool
	check = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}

		res := check(left.Left, right.Right)
		if res == false {
			return false
		}

		res = check(left.Right, right.Left)
		if res == false {
			return false
		}

		return true
	}

	return check(root.Left, root.Right)
}
