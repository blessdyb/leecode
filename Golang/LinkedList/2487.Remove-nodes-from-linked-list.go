package main


type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNodes(head *ListNode) *ListNode {
    stack := []*ListNode{}
    for head != nil {
        for len(stack) != 0 && head.Val > stack[len(stack) - 1].Val {
            stack = stack[:len(stack) - 1]
        }
        stack = append(stack, head)
        head = head.Next
    }
    stack[0].Next = nil
    for i := 1; i < len(stack); i++ {
        stack[i - 1].Next = stack[i]
        stack[i].Next = nil
    }
    return stack[0]
}