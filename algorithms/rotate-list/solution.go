package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	current := head
	count := 1
	for current.Next != nil {
		current = current.Next
		count += 1
	}
	current.Next = head

	k = k % count
	stepsToNewTail := count - k

	// Find the new Start
	newTail := current
	for i := 0; i < stepsToNewTail; i++ {
		newTail = newTail.Next
	}

	newHead := newTail.Next
	newTail.Next = nil

	return newHead
}
