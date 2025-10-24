package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// First solution, kinda bad tbh
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	vAns := []*TreeNode{}
	findLca(root, p, q, &vAns)
	return vAns[len(vAns)-1]
}

func findLca(root, p, q *TreeNode, vAns *[]*TreeNode) {
	if root == nil {
		return
	}

	pFound := false
	qFound := false
	checkLca(root, p, q, &pFound, &qFound)

	if pFound && qFound {
		*vAns = append(*vAns, root)
	}

	if root.Val != p.Val && root.Val != q.Val {
		findLca(root.Left, p, q, vAns)
		findLca(root.Right, p, q, vAns)
	}
}

func checkLca(root, p, q *TreeNode, pFound, qFound *bool) {
	if root == nil {
		return
	}

	if root.Val == p.Val {
		*pFound = true
	}
	if root.Val == q.Val {
		*qFound = true
	}

	checkLca(root.Left, p, q, pFound, qFound)
	checkLca(root.Right, p, q, pFound, qFound)
}

// Faster iteration
func lowestCommonAncestorAlternative(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	return right
}
