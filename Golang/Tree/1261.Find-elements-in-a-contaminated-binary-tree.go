package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type FindElements struct {
	root *TreeNode
}

func Constructor(root *TreeNode) FindElements {
	root.Val = 0
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			node.Left.Val = 2*node.Val + 1
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			node.Right.Val = 2*node.Val + 2
			stack = append(stack, node.Right)
		}
	}
	return FindElements{
		root: root,
	}
}

func (this *FindElements) Find(target int) bool {
	stack := []*TreeNode{this.root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		if node.Val == target {
			return true
		}
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	return false
}
