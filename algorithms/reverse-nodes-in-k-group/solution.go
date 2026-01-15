package main

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}

	res := &ListNode{}
	resPtn := res
	count := 1
	counter := head
	length := 0
	for counter != nil {
		length += 1
		counter = counter.Next
	}

	reversed := &ListNode{}
	reversedLast := &ListNode{}
	for head != nil {
		if count == 1 && length >= k {
			reversedLast = &ListNode{Val: head.Val}
			reversed = reversedLast

			head = head.Next
			count += 1
			length -= 1
			continue
		}

		if count < k && count > 1 && length > k-count {
			newFirst := &ListNode{Val: head.Val}
			tmp := reversed
			reversed = newFirst
			reversed.Next = tmp

			head = head.Next
			count += 1
			length -= 1
			continue
		}

		if count == k {
			newFirst := &ListNode{Val: head.Val}
			tmp := reversed
			reversed = newFirst
			reversed.Next = tmp

			res.Next = reversed
			res = reversedLast

			reversed = &ListNode{}
			reversedLast = &ListNode{}
			head = head.Next
			count = 1
			length -= 1
			continue
		}

		res.Next = &ListNode{Val: head.Val}
		res = res.Next

		count += 1
		length -= 1
		head = head.Next
	}

	return resPtn.Next
}
