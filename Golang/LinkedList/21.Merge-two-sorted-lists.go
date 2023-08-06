package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var ret, current *ListNode
	for list1 != nil && list2 != nil {
		var temp *ListNode
		if list1.Val < list2.Val {
			temp = &ListNode{
				Val: list1.Val,
			}
			list1 = list1.Next
		} else {
			temp = &ListNode{
				Val: list2.Val,
			}
			list2 = list2.Next
		}
		if ret == nil {
			ret = temp
			current = temp
		} else {
			current.Next = temp
			current = current.Next
		}
	}
	if list1 != nil {
		if current != nil {
			current.Next = list1
		} else {
			ret = list1
		}
	}
	if list2 != nil {
		if current != nil {
			current.Next = list2
		} else {
			ret = list2
		}
	}
	return ret
}
