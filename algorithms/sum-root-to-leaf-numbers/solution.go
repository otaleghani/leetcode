package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	tot := 0
	curr := []int{}
	var f func(r *TreeNode, sum *int)
	f = func(r *TreeNode, sum *int) {
		if r == nil {
			return
		}
		if r.Left == nil && r.Right == nil {
			curr = append(curr, r.Val)
			for i := 0; i < len(curr); i++ {
				v := curr[i]
				mult := len(curr) - i
				for range mult - 1 {
					v *= 10
				}
				*sum += v
			}
			// purge the last one
			curr = curr[:len(curr)-1]
			return
		}

		curr = append(curr, r.Val)
		f(r.Left, sum)
		f(r.Right, sum)
		curr = curr[:len(curr)-1]
	}

	f(root, &tot)
	return tot
}

// An easier to read approach
func sumNumbersAlt(root *TreeNode) int {
	return dfs(root, 0)
}
func dfs(node *TreeNode, currentSum int) int {
	if node == nil {
		return 0
	}

	currentSum = currentSum*10 + node.Val

	if node.Left == nil && node.Right == nil {
		return currentSum
	}

	return dfs(node.Left, currentSum) + dfs(node.Right, currentSum)
}
