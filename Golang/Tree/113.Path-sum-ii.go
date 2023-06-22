package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := [][]int{}
	if root != nil {
		type ComplextNode struct {
			Node    *TreeNode
			Parents []int
			target  int
		}
		stack := []ComplextNode{ComplextNode{
			Node:    root,
			Parents: []int{},
			target:  targetSum,
		}}
		for len(stack) > 0 {
			children := []ComplextNode{}
			for _, item := range stack {
				if item.Node.Val == item.target && item.Node.Left == nil && item.Node.Right == nil {
					ret = append(ret, append(item.Parents, item.Node.Val))
				}
				if item.Node.Left != nil {
					children = append(children, ComplextNode{
						Node:    item.Node.Left,
						Parents: append(item.Parents, item.Node.Val),
						target:  item.target - item.Node.Val,
					})
				}
				if item.Node.Right != nil {
					children = append(children, ComplextNode{
						Node:    item.Node.Right,
						Parents: append(item.Parents, item.Node.Val),
						target:  item.target - item.Node.Val,
					})
				}
			}
			stack = children
		}
	}
	return ret
}
