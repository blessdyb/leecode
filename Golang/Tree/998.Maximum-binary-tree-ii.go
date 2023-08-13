package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	if root != nil {
		if root.Val < val {
			return &TreeNode{
				Val:  val,
				Left: root,
			}
		} else if root.Right == nil {
			root.Right = &TreeNode{Val: val}
		} else if root.Right.Val < val {
			root.Right = &TreeNode{
				Val:  val,
				Left: root.Right,
			}
		} else {
			root.Right = insertIntoBST(root.Right, val)
		}
	}
	return root
}
