package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) int {
	count := 0
	if root == nil {
		return 0
	}

	pathSumRec(root, root.Val, targetSum, 0, &count)

	return count
}

func pathSumRec(root *TreeNode, currVal, targetSum, depth int, count *int) {
	if root == nil {
		return
	}

	if currVal == targetSum {
		*count += 1
	}

	if root.Left != nil {
		currValLeft := currVal + root.Left.Val
		pathSumRec(root.Left, currValLeft, targetSum, depth+1, count)
		if depth == 0 {
			pathSumRec(root.Left, root.Left.Val, targetSum, 0, count)
		}
	}

	if root.Right != nil {
		currValRight := currVal + root.Right.Val
		pathSumRec(root.Right, currValRight, targetSum, depth+1, count)
		if depth == 0 {
			pathSumRec(root.Right, root.Right.Val, targetSum, 0, count)
		}
	}
}

// Faster O(N) speed
func pathSumAlternative(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	prefixSums := make(map[int]int)
	prefixSums[0] = 1
	return dfs(root, 0, targetSum, prefixSums)
}

func dfs(node *TreeNode, currentPathSum int, targetSum int, prefixSums map[int]int) int {
	if node == nil {
		return 0
	}

	currentPathSum += node.Val

	count := prefixSums[currentPathSum-targetSum]

	prefixSums[currentPathSum]++

	count += dfs(node.Left, currentPathSum, targetSum, prefixSums)
	count += dfs(node.Right, currentPathSum, targetSum, prefixSums)

	prefixSums[currentPathSum]--

	return count
}
