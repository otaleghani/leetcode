package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	// Find center
	fast := head
	middle := head
	for fast != nil {
		middle = middle.Next
		fast = fast.Next.Next
	}

	// Reverse middle
	var tmp *ListNode
	for middle != nil {
		currTmp := middle.Next
		middle.Next = tmp
		tmp = middle
		middle = currTmp
	}
	middle = tmp

	// Find max
	max := 0
	for middle != nil {
		currentVal := head.Val + middle.Val
		if currentVal > max {
			max = currentVal
		}

		middle = middle.Next
		head = head.Next
	}

	return max
}
