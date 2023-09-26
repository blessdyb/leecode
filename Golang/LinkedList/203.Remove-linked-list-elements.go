package main

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}
	data := []int{}
	node := head
	for node != nil {
		if node.Val != val {
			data = append(data, node.Val)
		}
		node = node.Next
	}
	if len(data) == 0 {
		return nil
	}
	newHead := &ListNode{
		Val: data[0],
	}
	node = newHead
	for i := 1; i < len(data); i++ {
		node.Next = &ListNode{
			Val: data[i],
		}
		node = node.Next
	}
	return newHead
}
