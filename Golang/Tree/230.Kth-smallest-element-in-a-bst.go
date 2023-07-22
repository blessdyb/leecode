package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	current := root
	stack := []*TreeNode{}
	var ret int
	for k > 0 {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ret = node.Val
			k--
			current = node.Right
		} else {
			break
		}
	}
	return ret
}
