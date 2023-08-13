package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	// Edge case #1
	if depth == 1 {
		return &TreeNode{
			Val:  val,
			Left: root,
		}
	}
	stack := []*TreeNode{root}
	level := 1
	for len(stack) > 0 {
		child := []*TreeNode{}
		for _, s := range stack {
			if s.Left != nil {
				child = append(child, s.Left)
			}
			if s.Right != nil {
				child = append(child, s.Right)
			}
			if level+1 == depth {
				s.Left = &TreeNode{
					Val:  val,
					Left: s.Left,
				}
				s.Right = &TreeNode{
					Val:   val,
					Right: s.Right,
				}
			}
		}
		if level+1 == depth {
			break
		}
		level++
		stack = child
	}
	return root
}
