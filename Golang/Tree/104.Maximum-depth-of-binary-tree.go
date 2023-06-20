package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxDepthRecursive(root *TreeNode) int {
	var recursive func(node *TreeNode) int
	recursive = func(node *TreeNode) int {
		if node != nil {
			return 1 + max(recursive(node.Left), recursive(node.Right))
		}
		return 0
	}
	return recursive(root)
}

func maxDepthLayer(root *TreeNode) int {
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
			}
			ret += 1
			stack = children
		}
	}
	return ret
}
