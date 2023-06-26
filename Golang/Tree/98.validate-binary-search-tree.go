package main

import "math"

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Since the intuitive recursive solution couldn't guarteen any nodes on left can be smaller than the root node or any nodes on the right can be larger than the root nodes.
// For example: [5,4,6,null,null,3,7]. So if we go with the recursive solution, we need to pass-down the scope to the leaf nodes.
func isValidBSTRecursive(root *TreeNode) bool {
	var isValidBSTWithMinMax func(node *TreeNode, min, max int) bool
	isValidBSTWithMinMax = func(node *TreeNode, min, max int) bool {
		leftValid := isValidBSTWithMinMax(node.Left, min, node.Val)
		rightValid := isValidBSTWithMinMax(node.Right, node.Val, max)
		return leftValid && rightValid && node.Val > min && node.Val < max
	}
	return isValidBSTWithMinMax(root, math.MinInt64, math.MaxInt64)
}

// PreOrder BST results in a sorted list
func isValidBST(root *TreeNode) bool {
	if root != nil {
		previous := math.MinInt64
		node := root
		stack := []*TreeNode{}
		for {
			if node != nil {
				stack = append(stack, node)
				node = node.Left
			} else if len(stack) > 0 {
				temp := stack[len(stack)-1]
				stack = stack[:len(stack)-1]
				if temp.Val <= previous {
					return false
				}
				previous = temp.Val
				node = temp.Right
			} else {
				break
			}
		}
	}
	return true
}
