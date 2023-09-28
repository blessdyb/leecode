package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	current := node
	for current != nil && current.Next != nil {
		current.Val = current.Next.Val
		if current.Next.Next == nil {
			current.Next = nil
		} else {
			current = current.Next
		}
	}
}
