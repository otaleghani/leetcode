package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	tail := head.Next
	tailResult := &ListNode{}
	result := tailResult
	headResult := head

	for head.Next != nil {
		// Last element for the tail
		if head.Next.Next == nil {
			tailResult.Next = head.Next
			break
		}
		tail = head.Next      // tail: 2,3,4,5        // tail: 4,5
		head.Next = tail.Next // head: 1,3,4,5        // head: 3,5
		head = head.Next      // head: 3,4,5          // head: 5
		tail.Next = nil       // tail: 2              // tail: 4
		tailResult.Next = tail
		tailResult = tailResult.Next
	}

	head.Next = result.Next
	return headResult
}
