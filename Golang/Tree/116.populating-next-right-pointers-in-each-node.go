package main

/**
 * Definition for a binary tree node.
 */
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root != nil {
		stack := []*Node{root}
		for len(stack) > 0 {
			children := []*Node{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				if i < len(stack)-1 {
					node.Next = stack[i+1]
				}
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
			}
			stack = children
		}
	}
	return root
}
