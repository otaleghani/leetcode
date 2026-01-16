package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	target := &ListNode{}
	previous := &ListNode{}

	count := 0
	start := head

	for head != nil {
		count += 1
		head = head.Next
		if count == n {
			target = start
			previous = start
		}
		if count > n {
			previous = target
			target = target.Next
		}
	}

	if previous == start && target == start {
		return start.Next
	}
	previous.Next = target.Next

	return start
}
