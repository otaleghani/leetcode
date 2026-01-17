package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	res := &ListNode{}
	resPnt := res

	for head != nil {
		if head.Next != nil {
			if head.Val != head.Next.Val {
				res.Next = &ListNode{Val: head.Val}
				res = res.Next
				head = head.Next
			} else {
				curr := head.Val
				for head != nil && head.Val == curr {
					head = head.Next
				}
			}
		} else {
			res.Next = &ListNode{Val: head.Val}
			head = head.Next
		}
	}

	return resPnt.Next
}
