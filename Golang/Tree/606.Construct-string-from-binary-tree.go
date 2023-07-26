package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func tree2str(root *TreeNode) string {
	child := ""
	if root.Left != nil && root.Right != nil {
		child = "(" + tree2str(root.Left) + ")(" + tree2str(root.Right) + ")"
	} else if root.Left != nil {
		child = "(" + tree2str(root.Left) + ")"
	} else if root.Right != nil {
		child = "()(" + tree2str(root.Right) + ")"
	}
	return fmt.Sprintf("%v%v", root.Val, child)
}
