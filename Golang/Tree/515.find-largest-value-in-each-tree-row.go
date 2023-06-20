package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func largestValues(root *TreeNode) []int {
	ret := []int{}
	if root != nil {
		stack := []*TreeNode{root}
		for len(stack) > 0 {
			maxNum := math.MinInt32
			children := []*TreeNode{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				if maxNum < node.Val {
					maxNum = node.Val
				}
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
			}
			ret = append(ret, maxNum)
			stack = children
		}
	}
	return ret
}
