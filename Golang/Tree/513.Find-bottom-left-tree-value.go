package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findBottomLeftValue(root *TreeNode) int {
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		nextLevel := []*TreeNode{}
		for i := 0; i < len(stack); i++ {
			node := stack[i]
			if node.Left != nil {
				nextLevel = append(nextLevel, node.Left)
			}
			if node.Right != nil {
				nextLevel = append(nextLevel, node.Right)
			}
		}
		if len(nextLevel) == 0 {
			return stack[0].Val
		}
		stack = nextLevel
	}
	return -1
}
