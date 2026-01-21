package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {

	smaller := &ListNode{}
	larger := &ListNode{}
	start := smaller
	startLarger := larger

	for head != nil {
		if head.Val < x {
			smaller.Next = &ListNode{}
			smaller = smaller.Next
			smaller.Val = head.Val
		} else {
			larger.Next = &ListNode{}
			larger = larger.Next
			larger.Val = head.Val
		}
		head = head.Next
	}

	smaller.Next = startLarger.Next
	return start.Next
}
