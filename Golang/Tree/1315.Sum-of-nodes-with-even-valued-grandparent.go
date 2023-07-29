package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumEvenGrandparent(root *TreeNode) int {
	ret := 0
	stack := []*TreeNode{root}
	check := func(node *TreeNode, grandparent int) {
		if grandparent%2 == 0 {
			if node.Left != nil {
				ret += node.Left.Val
			}
			if node.Right != nil {
				ret += node.Right.Val
			}
		}
	}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			check(node.Left, node.Val)
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			check(node.Right, node.Val)
			stack = append(stack, node.Right)
		}
	}
	return ret
}
