package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func searchBSTRecursive(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	} else if root.Val == val {
		return root
	} else if root.Val > val {
		return searchBSTRecursive(root.Left, val)
	} else {
		return searchBSTRecursive(root.Right, val)
	}
}

func searchBSTPointer(root *TreeNode, val int) *TreeNode {
	node := root
	for node != nil {
		if node.Val == val {
			break
		} else if node.Val > val {
			node = node.Left
		} else {
			node = node.Right
		}
	}
	return node
}
