package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func increasingBST(root *TreeNode) *TreeNode {
	var newRoot, prev *TreeNode
	current := root
	stack := []*TreeNode{}
	for {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if newRoot == nil {
				newRoot = &TreeNode{
					Val: node.Val,
				}
				prev = newRoot
			} else {
				prev.Right = &TreeNode{
					Val: node.Val,
				}
				prev = prev.Right
			}
			current = node.Right
		} else {
			break
		}
	}
	return newRoot
}
