package main

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	parents := make(map[*TreeNode]*TreeNode)
	
	var buildParents func(node *TreeNode, parent *TreeNode)
	buildParents = func(node *TreeNode, parent *TreeNode) {
		if node == nil {
			return
		}
		parents[node] = parent
		buildParents(node.Left, node)
		buildParents(node.Right, node)
	}
	buildParents(root, nil)

	queue := []*TreeNode{target}
	visited := make(map[*TreeNode]bool)
	visited[target] = true
	dist := 0

	for len(queue) > 0 {
		if dist == k {
			var result []int
			for _, node := range queue {
				result = append(result, node.Val)
			}
			return result
		}
		
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[0]
			queue = queue[1:]

			if curr.Left != nil && !visited[curr.Left] {
				visited[curr.Left] = true
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil && !visited[curr.Right] {
				visited[curr.Right] = true
				queue = append(queue, curr.Right)
			}
			if parent, exists := parents[curr]; exists && parent != nil && !visited[parent] {
				visited[parent] = true
				queue = append(queue, parent)
			}
		}
		dist++
	}

	return []int{}
}
