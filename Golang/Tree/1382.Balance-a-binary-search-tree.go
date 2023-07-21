package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	data := []int{}
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			data = append(data, node.Val)
			dfs(node.Right)
		}
	}
	var helper func(nodes []int) *TreeNode
	helper = func(nodes []int) *TreeNode {
		var node *TreeNode
		if len(nodes) > 0 {
			mid := len(nodes) / 2
			node = &TreeNode{
				Val:   nodes[mid],
				Left:  helper(nodes[:mid]),
				Right: helper(nodes[mid+1:]),
			}
		}
		return node
	}
	dfs(root)
	return helper(data)
}
