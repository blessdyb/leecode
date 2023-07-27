package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTarget(root *TreeNode, k int) bool {
	hashmap := map[int]bool{}
	current := root
	stack := []*TreeNode{}
	for {
		if current != nil {
			stack = append(stack, current)
			current = current.Left
		} else if len(stack) > 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if hashmap[k-node.Val] {
				return true
			}
			hashmap[node.Val] = true
			current = node.Right
		} else {
			break
		}
	}
	return false
}
