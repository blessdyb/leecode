package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertGreatestCommonDivisors(head *ListNode) *ListNode {
	getMaxDivisor := func(a, b int) int {
		if a > b {
			a, b = b, a
		}
		for i := a; i >= 1; i-- {
			if a%i == 0 && b%i == 0 {
				return i
			}
		}
		return 1
	}
	node := head
	for node != nil {
		if node.Next != nil {
			num := getMaxDivisor(node.Val, node.Next.Val)
			next := node.Next
			temp := &ListNode{
				Val:  num,
				Next: next,
			}
			node.Next = temp
			node = next
		} else {
			node = node.Next
		}
	}
	return head
}
