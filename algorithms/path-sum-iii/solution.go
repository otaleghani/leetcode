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
