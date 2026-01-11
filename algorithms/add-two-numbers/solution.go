package main

// jDefinition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	carry := 0
	res := &ListNode{}
	original := res

	for l1 != nil || l2 != nil || carry != 0 {
		res.Next = &ListNode{}
		res = res.Next

		val1 := 0
		if l1 != nil {
			val1 = l1.Val
		}
		val2 := 0
		if l2 != nil {
			val2 = l2.Val
		}

		res.Val = val1 + val2 + carry
		if res.Val >= 10 {
			carry = 1
			res.Val -= 10
		} else {
			carry = 0
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return original.Next
}
