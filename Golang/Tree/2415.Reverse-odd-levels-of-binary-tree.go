package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func reverseOddLevels(root *TreeNode) *TreeNode {
	stack := []*TreeNode{root}
	level := 0
	for len(stack) > 0 {
		level++
		child := []*TreeNode{}
		for _, s := range stack {
			if s.Left != nil {
				child = append(child, s.Left, s.Right)
			}
		}
		if level%2 == 1 {
			for i, j := 0, len(child)-1; i < j; {
				child[i].Val, child[j].Val = child[j].Val, child[i].Val
				i++
				j--
			}
		}
		stack = child
	}
	return root
}
