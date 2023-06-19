package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root != nil {
		stack := []*TreeNode{root}
		length := 1
		for len(stack) > 0 {
			children := []*TreeNode{}
			level := []int{}
			for i := 0; i < length; i++ {
				node := stack[i]
				level = append(level, node.Val)
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
			}
			length = len(children)
			stack = children
			ret = append(ret, level)
		}
	}
	return ret
}
