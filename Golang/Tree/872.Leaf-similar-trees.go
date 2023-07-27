package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	getLeafNodes := func(root *TreeNode) []int {
		ret := []int{}
		current := root
		stack := []*TreeNode{}
		for {
			if current != nil {
				stack = append(stack, current)
				current = current.Left
			} else if len(stack) > 0 {
				node := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if node.Left == nil && node.Right == nil {
					ret = append(ret, node.Val)
				}
				current = node.Right
			} else {
				break
			}
		}
		return ret
	}
	root1Leaves := getLeafNodes(root1)
	root2Leaves := getLeafNodes(root2)
	if len(root1Leaves) == len(root2Leaves) {
		for i := 0; i < len(root1Leaves); i++ {
			if root1Leaves[i] != root2Leaves[i] {
				return false
			}
		}
		return true
	}
	return false
}
