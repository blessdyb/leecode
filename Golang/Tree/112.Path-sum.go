package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func hasPathSumRecursive(root *TreeNode, targetSum int) bool {
	var recursive func(node *TreeNode, value int) bool
	recursive = func(node *TreeNode, value int) bool {
		if node == nil {
			return false
		}
		if node.Val == value && node.Left == nil && node.Right == nil {
			return true
		}
		return recursive(node.Left, value-node.Val) || recursive(node.Right, value-node.Val)
	}
	return recursive(root, targetSum)
}

func hasPathSumBFS(root *TreeNode, targetSum int) bool {
	if root != nil {
		type ComplexNode struct {
			Node   *TreeNode
			Target int
		}
		stack := []ComplexNode{{
			Node:   root,
			Target: targetSum,
		}}
		for len(stack) > 0 {
			children := []ComplexNode{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				if node.Node.Val == node.Target && node.Node.Left == nil && node.Node.Right == nil {
					return true
				}
				if node.Node.Left != nil {
					children = append(children, ComplexNode{
						Node:   node.Node.Left,
						Target: node.Target - node.Node.Val,
					})
				}
				if node.Node.Right != nil {
					children = append(children, ComplexNode{
						Node:   node.Node.Right,
						Target: node.Target - node.Node.Val,
					})
				}
			}
			stack = children
		}
	}
	return false
}
