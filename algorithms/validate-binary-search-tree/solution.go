package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}

	valid := true
	var f func(*TreeNode, int, bool)
	f = func(r *TreeNode, depth int, isLeft bool) {
		if r == nil || !valid {
			return
		}

		rPnt := root
		currDepth := 0
		currIsLeft := false
		for rPnt != nil {
			if r.Val == rPnt.Val {
				break
			}
			if r.Val < rPnt.Val {
				currDepth += 1
				currIsLeft = true
				rPnt = rPnt.Left
			} else {
				currDepth += 1
				currIsLeft = false
				rPnt = rPnt.Right
			}
		}
		if rPnt == nil || currDepth != depth || isLeft != currIsLeft {
			valid = false
		}

		f(r.Left, depth+1, true)
		f(r.Right, depth+1, false)
	}
	f(root, 0, false)

	return valid
}
