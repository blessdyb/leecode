package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isEvenOddTree(root *TreeNode) bool {
	if root.Val%2 == 0 {
		return false
	}
	level := 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		children := []*TreeNode{}
		level++
		for _, s := range stack {
			if s.Left != nil {
				children = append(children, s.Left)
			}
			if s.Right != nil {
				children = append(children, s.Right)
			}
		}
		if len(children) > 0 {
			for i := 0; i < len(children); i++ {
				if level%2 == 0 && children[i].Val%2 == 0 {
					return false
				}
				if level%2 == 1 && children[i].Val%2 == 1 {
					return false
				}
				if i > 0 {
					if level%2 == 0 && children[i].Val <= children[i-1].Val {
						return false
					}
					if level%2 == 1 && children[i].Val >= children[i-1].Val {
						return false
					}
				}
			}
		}
		stack = children
	}
	return true
}
