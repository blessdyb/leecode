package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	ret := 0
	var helper func(node *TreeNode) int
	helper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := helper(node.Left)
		right := helper(node.Right)
		ret = max(ret, max(1+left+right, max(left, right)))
		return 1 + max(left, right)
	}
	helper(root)
	return ret - 1
}
