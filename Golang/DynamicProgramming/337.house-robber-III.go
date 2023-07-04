package main

import (
	"fmt"
	"strconv"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func robMemoryRecursive(root *TreeNode) int {
	cache := make(map[string]int)
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var recursive func(root *TreeNode, flag bool) int
	recursive = func(root *TreeNode, flag bool) int {
		key := fmt.Sprintf("%p_%s", root, strconv.FormatBool(flag))
		if value, key := cache[key]; key {
			return value
		}
		ret := 0
		if root != nil {
			if flag {
				ret = root.Val + recursive(root.Left, false) + recursive(root.Right, false)
			} else {
				ret = max(recursive(root.Left, true), recursive(root.Left, false)) + max(recursive(root.Right, true), recursive(root.Right, false))
			}
		}
		cache[key] = ret
		return ret
	}
	return max(recursive(root, true), recursive(root, false))
}

/*
 * Since we need to get all children nodes involved, it's better to make a post traverse. So we can make it done in O(n)
 */
func robDP(root *TreeNode) int {
	var max = func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	var postTraversal func(root *TreeNode) (int, int)
	postTraversal = func(root *TreeNode) (int, int) {
		if root == nil {
			return 0, 0
		}
		withLeft, withoutLeft := postTraversal(root.Left)
		withRight, withoutRight := postTraversal(root.Right)
		with := root.Val + withoutLeft + withoutRight
		without := max(withLeft, withoutLeft) + max(withRight, withoutRight)
		return with, without
	}
	return max(postTraversal(root))
}
