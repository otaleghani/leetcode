package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	og := &ListNode{}
	res := og

	for list1 != nil && list2 != nil {
		res.Next = &ListNode{}
		res = res.Next
		if list1.Val < list2.Val {
			res.Val = list1.Val
			list1 = list1.Next
		} else {
			res.Val = list2.Val
			list2 = list2.Next
		}
	}

	if list1 == nil && list2 != nil {
		res.Next = list2
	}
	if list2 == nil && list1 != nil {
		res.Next = list1
	}

	return og.Next
}
