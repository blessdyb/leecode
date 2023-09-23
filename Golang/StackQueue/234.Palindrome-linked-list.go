package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	ret := []int{}
	node := head
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	for i, j := 0, len(ret)-1; i < j; {
		if ret[i] != ret[j] {
			return false
		}
		i++
		j--
	}
	return true
}
