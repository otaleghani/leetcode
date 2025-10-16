package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// Really slow solution
func deleteMiddle(head *ListNode) *ListNode {
	list := []*ListNode{}

	for head != nil {
		list = append(list, head)
		head = head.Next
	}

	middleIndex := (len(list)) / 2
	if len(list) == 1 {
		return nil
	}
	if len(list) == 2 {
		list[len(list)-2].Next = nil
		return list[0]
	}

	list[middleIndex-1].Next = list[middleIndex+1]

	return list[0]
}

// Using two pointers far better
func deleteMiddleAlternative(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}

	dummy := &ListNode{Next: head}
	slow := dummy
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	slow.Next = slow.Next.Next

	return head
}
