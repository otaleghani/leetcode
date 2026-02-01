package main

// Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	Root  *TreeNode
	order []*TreeNode
	curr  int
}

func Constructor(root *TreeNode) BSTIterator {
	i := BSTIterator{
		Root: root,
	}

	order := []*TreeNode{}
	var f func(*TreeNode)
	f = func(root *TreeNode) {
		if root == nil {
			return
		}

		f(root.Left)
		order = append(order, root)
		f(root.Right)
	}
	f(root)
	i.order = order
	i.curr = 0
	return i
}

func (this *BSTIterator) Next() int {
	node := this.order[this.curr]
	this.curr += 1
	return node.Val
}

func (this *BSTIterator) HasNext() bool {
	if this.curr < len(this.order) {
		return true
	}

	return false
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
