package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	ret := 0
	var helper func(node *TreeNode, max int)
	helper = func(node *TreeNode, max int) {
		if node != nil {
			if node.Val >= max {
				ret++
				max = node.Val
			}
			helper(node.Left, max)
			helper(node.Right, max)
		}
	}
	helper(root, root.Val)
	return ret
}
