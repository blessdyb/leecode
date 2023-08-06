package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head != nil {
		prev := head
		current := head.Next
		for current != nil {
			if current.Val == prev.Val {
				current = current.Next
			} else {
				prev.Next = current
				prev = current
				current = current.Next
			}
		}
		// We must set the prev.Next to nil in case the last several items are the same
		prev.Next = nil
	}
	return head
}
