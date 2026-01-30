package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	items := make(map[int][]int)

	var f func(r *TreeNode, depth int)
	f = func(r *TreeNode, depth int) {
		if r == nil {
			return
		}

		if _, ok := items[depth]; !ok {
			items[depth] = []int{}
		}
		items[depth] = append(items[depth], r.Val)
		f(r.Left, depth+1)
		f(r.Right, depth+1)
	}
	f(root, 0)

	result := make([]float64, len(items))
	for depth, arr := range items {
		res := 0
		for _, item := range arr {
			res += item
		}

		result[depth] = float64(res) / float64(len(arr))
	}

	return result
}
