package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var getDepth func(node *TreeNode) int
	var checkBalanced func(node *TreeNode) bool
	getDepth = func(node *TreeNode) int {
		if node != nil {
			left := getDepth(node.Left)
			right := getDepth(node.Right)
			if left > right {
				return 1 + left
			}
			return 1 + right
		}
		return 0
	}
	checkBalanced = func(node *TreeNode) bool {
		if node != nil {
			left := getDepth(node.Left)
			right := getDepth(node.Right)
			if int(math.Abs(float64(left-right))) < 2 {
				return checkBalanced(node.Left) && checkBalanced(node.Right)
			}
			return false
		}
		return true
	}
	return checkBalanced(root)
}

func isBalancedOptimized(root *TreeNode) bool {
	var getDepth func(node *TreeNode) int
	getDepth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		left := getDepth(node.Left)
		right := getDepth(node.Right)
		if left == -1 || right == -1 {
			return -1
		}
		if int(math.Abs(float64(left-right))) > 1 {
			return -1
		}
		if left > right {
			return 1 + left
		}
		return 1 + right
	}
	return getDepth(root) != -1
}
