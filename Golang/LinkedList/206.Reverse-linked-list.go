package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	data := []int{}
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}
	head = &ListNode{
		Val: data[len(data)-1],
	}
	node := head
	for i := len(data) - 2; i >= 0; i-- {
		node.Next = &ListNode{
			Val: data[i],
		}
		node = node.Next
	}
	return head
}
