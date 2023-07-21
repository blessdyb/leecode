package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type TreeNodeWithParentValue struct {
	Node *TreeNode
	Val  int
}

func sumNumbers(root *TreeNode) int {
	stack := []*TreeNodeWithParentValue{&TreeNodeWithParentValue{
		Node: root,
		Val:  0,
	}}
	ret := 0
	for len(stack) > 0 {
		children := []*TreeNodeWithParentValue{}
		for i := 0; i < len(stack); i++ {
			node := stack[i]
			if node.Node.Left == nil && node.Node.Right == nil {
				ret += 10*node.Val + node.Node.Val
			}
			if node.Node.Left != nil {
				children = append(children, &TreeNodeWithParentValue{
					Node: node.Node.Left,
					Val:  10*node.Val + node.Node.Val,
				})
			}
			if node.Node.Right != nil {
				children = append(children, &TreeNodeWithParentValue{
					Node: node.Node.Right,
					Val:  10*node.Val + node.Node.Val,
				})
			}
		}
		stack = children
	}
	return ret
}
