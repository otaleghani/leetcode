package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var rec func(p []int, i []int) *TreeNode
	rec = func(p []int, i []int) *TreeNode {
		if len(p) == 0 {
			return nil
		}

		// Find the root
		rootVal := p[0]

		// Find the index of the root in the inorder
		rootInorderIndex := 0
		for idx, val := range i {
			if val == rootVal {
				rootInorderIndex = idx
			}
		}

		iLeft := i[:rootInorderIndex]
		iRight := i[rootInorderIndex+1:]

		pLeft := p[1 : 1+len(iLeft)]
		pRight := p[len(iLeft)+1:]

		curr := &TreeNode{Val: rootVal}
		curr.Left = rec(pLeft, iLeft)
		curr.Right = rec(pRight, iRight)

		return curr
	}

	return rec(preorder, inorder)
}
