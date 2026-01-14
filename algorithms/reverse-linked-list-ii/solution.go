package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {

	if left == right {
		return head
	}

	res := &ListNode{}
	resPnt := res
	pos := 1
	reversed := &ListNode{}
	reversedTail := &ListNode{}

	for head != nil {
		if pos == left {
			reversed.Val = head.Val
			pos += 1
			head = head.Next
			reversedTail = reversed
			continue
		} else if pos > left && pos < right {
			curr := &ListNode{
				Val: head.Val,
			}
			tmp := reversed
			reversed = curr
			curr.Next = tmp
			res.Next = reversed
			pos += 1
			head = head.Next
			continue
		} else if pos == right {
			curr := &ListNode{
				Val: head.Val,
			}
			tmp := reversed
			reversed = curr
			curr.Next = tmp
			res.Next = reversed
			res = reversedTail
			pos += 1
			head = head.Next
			continue
		}

		res.Next = &ListNode{}
		res = res.Next
		pos += 1
		res.Val = head.Val
		head = head.Next
	}
	return resPnt.Next
}
