package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func convertBST(root *TreeNode) *TreeNode {
	previous := 0
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node != nil {
			recursive(node.Right)
			node.Val += previous
			previous = node.Val
			recursive(node.Left)
		}
	}
	recursive(root)
	return root
}
