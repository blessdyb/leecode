package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodesRecursiveDFS(root *TreeNode) int {
	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node != nil {
			return 1 + recursive(node.Left) + recursive(node.Right)
		}
		return 0
	}
}

func countNodesRecursiveAlgebra(root *TreeNode) int {
	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node != nil {
			leftLevel, rightLevel := 0, 0
			current := node.Left
			for current != nil {
				current = current.Left
				leftLevel++
			}
			current = node.Right
			for current != nil {
				current = current.Right
				rightLevel++
			}
			if leftLevel == rightLevel {
				return (2 << leftLevel) - 1
			}
			return 1 + recursive(node.Left) + recursive(node.Right)
		}
		return 0
	}
	return recursive(root)
}
