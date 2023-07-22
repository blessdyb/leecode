package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	var dfs func(node *TreeNode) *TreeNode
	dfs = func(node *TreeNode) *TreeNode {
		if node.Left != nil {
			node.Left = dfs(node.Left)
		}
		if node.Right != nil {
			node.Right = dfs(node.Right)
		}
		if node.Val == target && node.Left == nil && node.Right == nil {
			return nil
		}
		return node
	}
	return dfs(root)
}
