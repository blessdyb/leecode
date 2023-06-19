package main

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	ret := [][]int{}
	if root != nil {
		stack := []*Node{root}
		for len(stack) > 0 {
			layer := []int{}
			children := []*Node{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				layer = append(layer, node.Val)
				children = append(children, node.Children...)
			}
			ret = append(ret, layer)
			stack = children
		}
	}
	return ret
}
