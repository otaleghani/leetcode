package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	if key > root.Val {
		root.Right = deleteNode(root.Right, key)
	} else if key < root.Val {
		root.Left = deleteNode(root.Left, key)
	} else {
		if root.Left == nil {
			tmp := root.Right
			root = nil
			return tmp
		}
		if root.Right == nil {
			tmp := root.Left
			root = nil
			return tmp
		}

		succ := getSuccessor(root)
		root.Val = succ.Val
		root.Right = deleteNode(root.Right, succ.Val)
	}

	return root
}

func getSuccessor(root *TreeNode) *TreeNode {
	root = root.Right
	for root != nil && root.Left != nil {
		root = root.Left
	}
	return root
}
