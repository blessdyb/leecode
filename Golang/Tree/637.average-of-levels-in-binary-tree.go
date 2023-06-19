package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfLevels(root *TreeNode) []float64 {
	ret := []float64{}
	if root != nil {
		stack := []*TreeNode{root}
		for len(stack) > 0 {
			children := []*TreeNode{}
			sum := 0
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				sum += node.Val
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
			}
			ret = append(ret, float64(sum)/float64(len(stack)))
			stack = children
		}
	}
	return ret
}
