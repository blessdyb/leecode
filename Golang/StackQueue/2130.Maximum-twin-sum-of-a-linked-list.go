package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	ret := []int{}
	node := head
	for node != nil {
		ret = append(ret, node.Val)
		node = node.Next
	}
	max := 0
	for i, j := 0, len(ret)-1; i < j; {
		if max < ret[i]+ret[j] {
			max = ret[i] + ret[j]
		}
		i++
		j--
	}
	return max
}
