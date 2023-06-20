package main

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	var recursive func(node *Node) int
	recursive = func(node *Node) int {
		if node != nil {
			maxValue := 0
			for _, n := range node.Children {
				if currentDepth := recursive(n); currentDepth > maxValue {
					maxValue = currentDepth
				}
			}
			return maxValue + 1
		}
		return 0
	}
	return recursive(root)
}

func maxDepthBFS(root *Node) int {
	ret := 0
	if root != nil {
		stack := []*Node{root}
		for len(stack) > 0 {
			children := []*Node{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				children = append(children, node.Children...)
			}
			ret++
			stack = children
		}
	}
	return ret
}
