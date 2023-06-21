package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func binaryTreePathsRecursive(root *TreeNode) []string {
	ret := []string{}
	if root != nil {
		var recursive func(node *TreeNode, parent string)
		recursive = func(node *TreeNode, parent string) {
			nextParent := strconv.Itoa(node.Val)
			if parent != "" {
				nextParent = parent + "->" + nextParent
			}
			if node.Left != nil {
				recursive(node.Left, nextParent)
			}
			if node.Right != nil {
				recursive(node.Right, nextParent)
			}
			if node.Left == nil && node.Right == nil {
				ret = append(ret, nextParent)
			}
		}
		recursive(root, "")
	}
	return ret
}

func binaryTreePathsBFS(root *TreeNode) []string {
	ret := []string{}
	if root != nil {
		type ComplexNode struct {
			Node   *TreeNode
			Parent string
		}
		stack := []ComplexNode{{
			Node:   root,
			Parent: "",
		}}
		for len(stack) > 0 {
			children := []ComplexNode{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				nextParent := strconv.Itoa(node.Node.Val)
				if node.Parent != "" {
					nextParent = node.Parent + "->" + nextParent
				}
				if node.Node.Left != nil {
					children = append(children, ComplexNode{
						Node:   node.Node.Left,
						Parent: nextParent,
					})
				}
				if node.Node.Right != nil {
					children = append(children, ComplexNode{
						Node:   node.Node.Right,
						Parent: nextParent,
					})
				}
				if node.Node.Left == nil && node.Node.Right == nil {
					ret = append(ret, nextParent)
				}
			}
			stack = children
		}
	}
	return ret
}
