package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	head := &ListNode{}
	current := head
	for list1 != nil && a > 0 {
		current.Next = &ListNode{Val: list1.Val}
		current = current.Next
		a--
		b--
		list1 = list1.Next
	}
	for list2 != nil {
		current.Next = &ListNode{Val: list2.Val}
		current = current.Next
		list2 = list2.Next
	}
	for list1 != nil && b >= 0 {
		b--
		list1 = list1.Next
	}
	for list1 != nil {
		current.Next = &ListNode{Val: list1.Val}
		current = current.Next
		list1 = list1.Next
	}

	return head.Next
}
