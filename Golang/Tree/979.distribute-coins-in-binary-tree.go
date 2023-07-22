package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Post order traversal
func distributeCoins(root *TreeNode) int {
	ret := 0
	abs := func(value int) int {
		return int(math.Abs(float64(value)))
	}
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		ret += abs(left) + abs(right)
		return node.Val + left + right - 1
	}
	return ret
}
