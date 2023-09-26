package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	var startNode *ListNode
	pointer := head
	sum := 0
	for pointer != nil {
		if pointer.Val == 0 {
			if startNode != nil {
				temp := &ListNode{Val: sum}
				startNode.Next = temp
				temp.Next = pointer
				sum = 0
			}
			startNode = pointer
		} else {
			sum += pointer.Val
		}
		pointer = pointer.Next
	}
	pointer = head.Next
	for pointer != nil && pointer.Next != nil {
		pointer.Next = pointer.Next.Next
		pointer = pointer.Next
	}
	return head.Next
}
