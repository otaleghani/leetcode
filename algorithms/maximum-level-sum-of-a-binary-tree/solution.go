package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	m := make(map[int]int)

	sumLevel(root, 1, &m)

	max := -99999
	result := 0
	for key, val := range m {
		if val > max {
			max = val
			result = key
		}
	}

	return result
}

func sumLevel(root *TreeNode, depth int, m *map[int]int) {
	if root == nil {
		return
	}

	(*m)[depth] += root.Val
	sumLevel(root.Right, depth+1, m)
	sumLevel(root.Left, depth+1, m)
}
