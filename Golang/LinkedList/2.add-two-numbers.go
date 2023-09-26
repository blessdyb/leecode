package main

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{}
	prev := head
	carrier := 0
	for l1 != nil || l2 != nil {
		v1, v2 := 0, 0
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			v1 = l2.Val
			l2 = l2.Next
		}
		value := v1 + v2 + carrier
		if value > 9 {
			value -= 10
			carrier = 1
		} else {
			carrier = 0
		}
		prev.Next = &ListNode{
			Val: value,
		}
		prev = prev.Next
	}
	if carrier == 1 {
		perv.Next = &ListNode{Val: 1}
	}
	return head.Next
}
