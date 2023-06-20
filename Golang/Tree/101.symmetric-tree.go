package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	var recursive func(n1, n2 *TreeNode) bool
	recursive = func(n1, n2 *TreeNode) bool {
		if (n1 != nil && n2 == nil) || (n1 == nil && n2 != nil) {
			return false
		} else if n1 != nil && n2 != nil {
			if n1.Val == n2.Val {
				return recursive(n1.Left, n2.Right) && recursive(n1.Right, n2.Left)
			}
			return false
		}
		return true
	}
	return recursive(root.Left, root.Right)
}

func isSymmetricBFS(root *TreeNode) bool {
	if root != nil {
		stack := []*TreeNode{root.Left, root.Right}
		for len(stack) > 0 {
			children := []*TreeNode{}
			for len(stack) > 0 {
				n1, n2 := stack[0], stack[1]
				stack = stack[2:]
				if (n1 != nil && n2 == nil) || (n1 == nil && n2 != nil) {
					return false
				} else if n1 == nil && n2 == nil {
					continue
				} else if n1.Val != n2.Val {
					return false
				}
				children = append(children, n1.Left, n2.Right, n1.Right, n2.Left)
			}
			stack = children
		}
	}
	return true
}
