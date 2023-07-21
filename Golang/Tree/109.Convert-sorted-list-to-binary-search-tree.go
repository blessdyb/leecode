package main

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	var helper func(nodes []int) *TreeNode
	helper = func(nodes []int) *TreeNode {
		var node *TreeNode
		if len(nodes) > 0 {
			mid := len(nodes) / 2
			node = &TreeNode{
				Val:   nodes[mid],
				Left:  helper(nodes[:mid]),
				Right: helper(nodes[mid+1:]),
			}
		}
		return node
	}
	data := []int{}
	for head != nil {
		data = append(data, head.Val)
		head = head.Next
	}

	return helper(data)
}
