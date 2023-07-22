package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	ret := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		children := []*TreeNode{}
		for _, item := range stack {
			if item.Left != nil {
				children = append(children, item.Left)
			}
			if item.Right != nil {
				children = append(children, item.Right)
			}
		}
		if len(children) == 0 {
			for _, item := range stack {
				ret += item.Val
			}
		}
		stack = children
	}
	return ret
}
