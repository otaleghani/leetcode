package main

import "math"

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var getGain func(*TreeNode) int
	getGain = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		// Ignore negative
		leftGain := max(0, getGain(node.Left))
		rightGain := max(0, getGain(node.Right))
		currPathSum := node.Val + leftGain + rightGain
		if currPathSum > maxSum {
			maxSum = currPathSum
		}
		return node.Val + max(leftGain, rightGain)
	}

	getGain(root)
	return maxSum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
