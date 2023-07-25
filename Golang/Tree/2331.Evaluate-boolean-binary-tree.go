package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func evaluateTree(root *TreeNode) bool {
	var helper func(node *TreeNode) bool
	helper = func(node *TreeNode) bool {
		if node.Left == nil && node.Right == nil {
			return node.Val == 1
		} else if node.Val == 2 {
			return helper(node.Left) || helper(node.Right)
		} else if node.Val == 3 {
			return helper(node.Left) && helper(node.Right)
		}
		return true
	}
	return helper(root)
}
