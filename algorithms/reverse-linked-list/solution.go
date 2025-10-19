package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var tmp *ListNode

	for head.Next != nil {
		tmpNext := head.Next // Preserve the next
		head.Next = tmp      // Reset the next with the value before
		tmp = head
		head = tmpNext
	}

	head.Next = tmp

	return head
}
