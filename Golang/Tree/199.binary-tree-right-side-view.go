package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	ret := []int{}
	if root != nil {
		stack := []*TreeNode{root}
		for len(stack) > 0 {
			layer := []*TreeNode{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				if node.Left != nil {
					layer = append(layer, node.Left)
				}
				if node.Right != nil {
					layer = append(layer, node.Right)
				}
			}
			ret = append(ret, stack[len(stack)-1].Val)
			stack = layer
		}
	}
	return ret
}
