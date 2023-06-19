package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversalRecursive(root *TreeNode) []int {
	ret := []int{}
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node != nil {
			recursive(node.Left)
			recursive(node.Right)
			ret = append(ret, node.Val)
		}
	}
	recursive(root)
	return ret
}

func postorderTraversalPointer(root *TreeNode) []int {
	ret := []int{}
	if root != nil {

		stack := []*TreeNode{root}
		for len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if node.Left != nil {
				stack = append(stack, node.Left)
			}
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			ret = append(ret, node.Val)
		}
	}
	return ret
}
