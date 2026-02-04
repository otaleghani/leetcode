package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	order := [][]int{}

	var f func(*TreeNode, int)
	f = func(root *TreeNode, level int) {
		if root == nil {
			return
		}

		if len(order)-1 < level {
			order = append(order, []int{})
		}

		order[level] = append(order[level], root.Val)
		f(root.Right, level+1)
		f(root.Left, level+1)
	}
	f(root, 0)

	for i, s := range order {
		if i%2 == 0 {
			for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
				s[i], s[j] = s[j], s[i]
			}
		}
	}

	return order
}
