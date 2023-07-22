package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findSecondMinimumValue(root *TreeNode) int {
	if root.Left != nil {
		stack := []*TreeNode{root.Left, root.Right}
		cursor := math.MaxInt32 + 1
		for len(stack) > 0 {
			node := stack[0]
			stack = stack[1:]
			if node.Val != root.Val && node.Val < cursor {
				cursor = node.Val
			}
			if node.Left != nil {
				stack = append(stack, node.Left, node.Right)
			}
		}
		if cursor != math.MaxInt32+1 {
			return cursor
		}
	}
	return -1
}
