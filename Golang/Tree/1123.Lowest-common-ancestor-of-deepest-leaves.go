package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	type TreeNodeWithParent struct {
		Node   *TreeNode
		Parent *TreeNodeWithParent
	}
	stack := []*TreeNodeWithParent{&TreeNodeWithParent{
		Node:   root,
		Parent: nil,
	}}
	for len(stack) > 0 {
		child := []*TreeNodeWithParent{}
		for _, s := range stack {
			if s.Node.Left != nil {
				child = append(child, &TreeNodeWithParent{
					Node:   s.Node.Left,
					Parent: s,
				})
			}
			if s.Node.Right != nil {
				child = append(child, &TreeNodeWithParent{
					Node:   s.Node.Right,
					Parent: s,
				})
			}
		}
		if len(child) == 0 {
			break
		}
		stack = child
	}
	if len(stack) == 1 {
		return stack[0].Node
	}
	for len(stack) > 1 {
		allEquals := true
		for i := 1; i < len(stack); i++ {
			if stack[0].Parent != stack[i].Parent {
				allEquals = false
				break
			}
		}
		if allEquals {
			stack = stack[:1]
		} else {
			for i := 0; i < len(stack); i++ {
				stack[i] = stack[i].Parent
			}
		}
	}
	return stack[0].Parent.Node
}
