package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func nextLargerNodes(head *ListNode) []int {
	values := []int{}
	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}
	ret := make([]int, len(values))
	stack := []int{0}
	for i := 1; i < len(values); i++ {
		for len(stack) > 0 && values[stack[len(stack)-1]] < values[i] {
			ret[stack[len(stack)-1]] = values[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	return ret
}
