package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pruneTree(root *TreeNode) *TreeNode {
	var helper func(node *TreeNode) *TreeNode
	helper = func(node *TreeNode) *TreeNode {
		if node != nil {
			node.Left = helper(node.Left)
			node.Right = helper(node.Right)
			if node.Val == 1 || node.Left != nil || node.Right != nil {
				return node
			}
			return nil
		}
		return nil
	}
	return helper(root)
}
