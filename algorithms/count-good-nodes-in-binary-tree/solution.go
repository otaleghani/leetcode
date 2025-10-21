package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return findGoodNodes(root, root.Val) / 2
}

// max is the original value of the root
func findGoodNodes(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	countLeft := 0
	countRight := 0

	if root.Val >= max {
		countLeft = 1 + findGoodNodes(root.Left, root.Val)
		countRight = 1 + findGoodNodes(root.Right, root.Val)
	} else {
		countLeft = findGoodNodes(root.Left, max)
		countRight = findGoodNodes(root.Right, max)
	}

	return countLeft + countRight
}

func goodNodesAlternative(root *TreeNode) int {
	count := 0
	findGoodNodesAlternative(root, root.Val, &count)

	return count
}

// max is the original value of the root
func findGoodNodesAlternative(root *TreeNode, max int, count *int) {
	if root == nil {
		return
	}

	if root.Val >= max {
		*count++
		max = root.Val
	}

	findGoodNodesAlternative(root.Left, max, count)
	findGoodNodesAlternative(root.Right, max, count)
}
