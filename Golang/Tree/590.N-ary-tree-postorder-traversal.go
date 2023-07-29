package main

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	ret := []int{}
	var helper func(node *Node)
	helper = func(node *Node) {
		if node != nil {
			for _, c := range node.Children {
				helper(c)
			}
			ret = append(ret, node.Val)
		}
	}
	helper(root)
	return ret
}
