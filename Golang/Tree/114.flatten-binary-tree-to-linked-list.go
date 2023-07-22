package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root != nil {
		if root.Left != nil {
			prev := root.Left
			for prev.Right != nil {
				prev = prev.Right
			}
			prev.Right = root.Right
			root.Right = root.Left
			root.Left = nil
		}
		flatten(root.Right)
	}
}
