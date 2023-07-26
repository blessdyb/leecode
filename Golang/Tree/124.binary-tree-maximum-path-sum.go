package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxPathSum(root *TreeNode) int {
	ret := 3 * 10000 * -1000
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := max(helper(node.Left), 0)
		right := max(helper(node.Right), 0)
		// the max from left => node => right
		ret = max(ret, node.Val+left+right)
		// the value goes through node and left/right for upper level
		return node.Val + max(left, right)
	}
	helper(root)
	return ret
}
