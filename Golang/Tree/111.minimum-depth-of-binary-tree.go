package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func minDepthRecursive(root *TreeNode) int {
	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node != nil {
			// Note: we need to reach to the leaf nodes (both left/right nodes are nil)
			if node.Left == nil {
				return 1 + recursive(node.Right)
			} else if node.Right == nil {
				return 1 + recursive(node.Left)
			}
			return 1 + min(recursive(node.Left), recursive(node.Right))
		}
		return 0
	}
	return recursive(root)
}

func minDepthLayer(root *TreeNode) int {
	ret := 0
	if root != nil {
		stack := []*TreeNode{root}
		for len(stack) > 0 {
			children := []*TreeNode{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
				if node.Left == nil && node.Right == nil {
					return 1 + ret
				}
			}
			ret += 1
			stack = children
		}
	}
	return ret
}
