package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func longestZigZag(root *TreeNode) int {
	if root == nil {
		return 0
	}

	maxLength := 0
	dfs(root, 0, &maxLength, true)

	return maxLength
}

func dfs(root *TreeNode, length int, maxLength *int, isLeft bool) {
	if root == nil {
		return
	}

	*maxLength = max(*maxLength, length)

	if isLeft {
		dfs(root.Left, length+1, maxLength, false)
		dfs(root.Right, 1, maxLength, true)
	} else {
		dfs(root.Right, length+1, maxLength, true)
		dfs(root.Left, 1, maxLength, false)
	}
}
