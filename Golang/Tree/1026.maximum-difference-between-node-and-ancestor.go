package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	ret := 0
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	min := func(a, b int) int {
		if a > b {
			return b
		}
		return a
	}
	var helper func(node *TreeNode, maxVal, minVal int)
	helper = func(node *TreeNode, maxVal, minVal int) {
		if node != nil {
			maxVal = max(node.Val, maxVal)
			minVal = min(node.Val, minVal)
			ret = max(ret, max(int(math.Abs(float64(node.Val-maxVal))), int(math.Abs(float64(node.Val-minVal)))))
			helper(node.Left, maxVal, minVal)
			helper(node.Right, maxVal, minVal)
		}
	}
	helper(root, root.Val, root.Val)
	return ret
}
