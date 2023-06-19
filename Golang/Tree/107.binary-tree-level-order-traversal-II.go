package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	ret := [][]int{}
	if root != nil {
		stack := []*TreeNode{root}
		layer := 1
		for len(stack) > 0 {
			children := []*TreeNode{}
			temp := []int{}
			for i := 0; i < layer; i++ {
				node := stack[i]
				temp = append(temp, node.Val)
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
			}
			ret = append(ret, temp)
			layer = len(children)
			stack = children
		}
	}
	for i, j := 0, len(ret)-1; i <= j; {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}
	return ret
}
