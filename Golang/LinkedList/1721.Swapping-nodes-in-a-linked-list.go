package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapNodes(head *ListNode, k int) *ListNode {
	count := 0
	current := head
	start, end := head, head
	for current != nil {
		count++
		if count == k {
			start = current
		}
		current = current.Next
	}
	target := count - k
	for end != nil && target != 0 {
		target--
		end = end.Next
	}
	start.Val, end.Val = end.Val, start.Val
	return head
}
