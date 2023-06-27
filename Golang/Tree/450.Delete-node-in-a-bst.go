package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root != nil {
		if root.Val == key {
			if root.Right == nil {
				return root.Left
			}
			node := root.Right
			for node != nil && node.Left != nil {
				node = node.Left
			}
			node.Left = root.Left
			return root.Right
		} else if root.Val < key {
			root.Right = deleteNode(root.Right, key)
		} else {
			root.Left = deleteNode(root.Left, key)
		}
	}
	return root
}
