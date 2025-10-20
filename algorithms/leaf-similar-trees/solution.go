package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	left := leafValueSequence(root1)
	right := leafValueSequence(root2)

	if len(left) != len(right) {
		return false
	}
	for i := 0; i < len(left); i++ {
		if left[i] != right[i] {
			return false
		}
	}

	return true
}

func leafValueSequence(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	left := leafValueSequence(root.Left)
	right := leafValueSequence(root.Right)

	if root.Left == nil && root.Right == nil {
		left = append(left, root.Val)
		right = append(right, root.Val)
	}

	return append(left, right...)
}
