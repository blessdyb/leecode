package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func getTargetCopy(original, cloned, target *TreeNode) *TreeNode {
	current := cloned
	stack := []*TreeNode{}
	for {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			current = node.Right
			if node.Val == target.Val {
				return node
			}
		} else {
			break
		}
	}
	return current
}
