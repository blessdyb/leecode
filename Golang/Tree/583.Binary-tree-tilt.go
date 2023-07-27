package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	ret := 0
	var dfs func(node *TreeNode) int
	dfs = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := dfs(node.Left)
		right := dfs(node.Right)
		tilt := left - right
		if tilt < 0 {
			tilt *= -1
		}
		ret += tilt
		return node.Val + left + right
	}
	dfs(root)
	return ret
}
