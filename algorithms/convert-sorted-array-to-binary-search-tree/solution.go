package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	// Divide it in two
	// Start from the end and go back up

	var f func([]int) *TreeNode
	f = func(arr []int) *TreeNode {
		if len(arr) == 0 {
			return nil
		}
		if len(arr) == 1 {
			return &TreeNode{Val: arr[0]}
		}

		// Find center
		center := len(arr) / 2

		return &TreeNode{
			Val:   arr[center],
			Left:  f(arr[:center]),
			Right: f(arr[center+1:]),
		}
	}

	return f(nums)
}
