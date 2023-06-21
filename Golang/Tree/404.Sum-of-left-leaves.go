package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumOfLeftLeavesRecursive(root *TreeNode) int {
	ret := 0
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node != nil {
			if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
				ret += node.Left.Val
			}
			recursive(node.Left)
			recursive(node.Right)
		}
	}
	recursive(root)
	return ret
}

func sumOfLeftLeavesStack(root *TreeNode) int {
	ret := 0
	if root != nil {
		stack := []*TreeNode{}
		node := root
		for {
			if node != nil {
				stack = append(stack, node)
				node = node.Left
			} else if len(stack) > 0 {
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if node.Left != nil && node.Left.Left == nil && node.Left.Right == nil {
					ret += node.Left.Val
				}
				node = node.Right
			} else {
				break
			}
		}
	}
	return ret
}
