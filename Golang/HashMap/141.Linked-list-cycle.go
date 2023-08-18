package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	visited := map[*ListNode]bool{}
	for head != nil {
		if visited[head] {
			return true
		}
		visited[head] = true
		head = head.Next
	}
	return false
}
