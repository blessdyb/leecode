package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) [][]int {
	ret := [][]int{}
	if root != nil {
		stack := []*TreeNode{root}
		ltr := true
		for len(stack) > 0 {
			children := []*TreeNode{}
			temp := []int{}
			for i := 0; i < len(stack); i++ {
				node := stack[i]
				temp = append(temp, node.Val)
				if node.Left != nil {
					children = append(children, node.Left)
				}
				if node.Right != nil {
					children = append(children, node.Right)
				}
			}
			stack = children
			if !ltr {
				for i, j := 0, len(temp)-1; i < j; {
					temp[i], temp[j] = temp[j], temp[i]
					i++
					j--
				}
			}
			ret = append(ret, temp)
			ltr = !ltr
		}
	}
	return ret
}
