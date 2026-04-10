package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var stack1, stack2 []int

	for l1 != nil {
		stack1 = append(stack1, l1.Val)
		l1 = l1.Next
	}

	for l2 != nil {
		stack2 = append(stack2, l2.Val)
		l2 = l2.Next
	}

	var head *ListNode
	carry := 0

	for len(stack1) > 0 || len(stack2) > 0 || carry > 0 {
		sum := carry

		if len(stack1) > 0 {
			sum += stack1[len(stack1)-1]
			stack1 = stack1[:len(stack1)-1]
		}

		if len(stack2) > 0 {
			sum += stack2[len(stack2)-1]
			stack2 = stack2[:len(stack2)-1]
		}

		carry = sum / 10
		newNode := &ListNode{
			Val:  sum % 10,
			Next: head,
		}

		head = newNode
	}

	return head
}
