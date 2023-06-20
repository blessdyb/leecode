package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node != nil {
			recursive(node.Left)
			recursive(node.Right)
			node.Left, node.Right = node.Right, node.Left
		}
	}
	recursive(root)
	return root
}
