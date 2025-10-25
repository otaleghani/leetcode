package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	m := make([]int, 0, 100)
	m = append(m, root.Val)
	findView(root, 0, &m)

	return m
}

func findView(root *TreeNode, depth int, m *[]int) {
	if root == nil {
		return
	}

	if root.Right != nil {
		if len(*m) == depth+1 {
			*m = append(*m, root.Right.Val)
			findView(root.Right, depth+1, m)
		} else {
			findView(root.Right, depth+1, m)
		}
	}
	if root.Left != nil {
		if len(*m) == depth+1 {
			*m = append(*m, root.Left.Val)
			findView(root.Left, depth+1, m)
		} else {
			findView(root.Left, depth+1, m)
		}
	}
}
