package main

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversalRecursive(root *TreeNode) []int {
	var ret []int
	var recursive func(node *TreeNode)
	recursive = func(node *TreeNode) {
		if node != nil {
			recursive(node.Left)
			ret = append(ret, node.Val)
			recursive(node.Right)
		}
	}
	return ret
}

func inorderTraversalPointer(root *TreeNode) []int {
	ret := []int{}
	stack := []*TreeNode{}
	node := root
	for {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else if len(stack) > 0 {
			temp := stack[len(stack)-1]
			ret = append(ret, temp.Val)
			stack = stack[:len(stack)-1]
			node = temp.Right
		} else {
			break
		}
	}
	return ret
}

func inorderTraversalMorris(root *TreeNode) []int {
	ret := []int{}
	node := root
	if node != nil {
		if node.Left == nil {
			ret = append(ret, node.Val)
			node = node.Right
		} else {
			previous := node.Left
			for previous.Right != nil && previous.Right != node {
				previous = previous.Right
			}
			if previous.Right == nil {
				previous.Right = node
				node = node.Left
			}
			if previous.Right == node {
				previous.Right = nil
				ret = append(ret, node.Val)
				node = node.Right
			}

		}
	}
	return ret
}
