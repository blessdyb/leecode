package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxLevelSum(root *TreeNode) int {
	max, level := root.Val, 1
	ret := level
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		child := []*TreeNode{}
		temp := 0
		for _, s := range stack {
			if s.Left != nil {
				temp += s.Left.Val
				child = append(child, s.Left)
			}
			if s.Right != nil {
				temp += s.Right.Val
				child = append(child, s.Right)
			}
		}
		level++
		// Edge case if it's the last layer
		if len(child) > 0 {
			if max < temp {
				max = temp
				ret = level
			}
		}

		stack = child
	}
	return ret
}
