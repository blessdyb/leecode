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
// For example: [5,4,6,null,null,3,7](https://camo.githubusercontent.com/2b2f3322ff7c6a9a30251309afba36172e1d1d74f4b917026a4451173f47dfe5/68747470733a2f2f636f64652d7468696e6b696e672d313235333835353039332e66696c652e6d7971636c6f75642e636f6d2f706963732f32303233303331303030303832342e706e67). So if we go with the recursive solution, we need to pass-down the scope to the leaf nodes.
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
