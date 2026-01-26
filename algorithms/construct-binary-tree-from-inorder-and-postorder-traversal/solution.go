package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	val := postorder[len(postorder)-1]

	// Find index in inorder
	indexValInorder := 0
	for i, v := range inorder {
		if v == val {
			indexValInorder = i
		}
	}

	lIn := inorder[:indexValInorder]
	rIn := inorder[indexValInorder+1:]

	rPo := postorder[len(lIn) : len(postorder)-1]
	lPo := postorder[:len(lIn)]

	left := buildTree(lIn, lPo)
	right := buildTree(rIn, rPo)

	return &TreeNode{Val: val, Left: left, Right: right}
}
