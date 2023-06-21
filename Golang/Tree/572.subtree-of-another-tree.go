package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtreeRecursive(root *TreeNode, subRoot *TreeNode) bool {
	var isSameTree func(r, s *TreeNode) bool
	var isSubTreeRecursive func(r, s *TreeNode) bool
	isSameTree = func(r, s *TreeNode) bool {
		if (r != nil && s == nil) || (r == nil && s != nil) {
			return false
		} else if r == nil && s == nil {
			return true
		} else if r.Val != s.Val {
			return false
		}
		return isSameTree(r.Left, s.Left) && isSameTree(r.Right, s.Right)
	}
	isSubTreeRecursive = func(r, s *TreeNode) bool {
		if !isSameTree(r, s) {
			if r != nil {
				return isSubTreeRecursive(r.Left, s) || isSubTreeRecursive(r.Right, s)
			}
			return false
		}
		return true
	}
	return isSubTreeRecursive(root, subRoot)
}

// It turns out that if we include # or any other character for the null node while serializing, then we can uniquely identify a tree, that too with only one traversal (either preorder or postorder).
func isSubtreeIndex(root *TreeNode, subRoot *TreeNode) bool {
	var traversal func(node *TreeNode, flag string) string
	traversal = func(node *TreeNode, flag string) string {
		if node != nil {
			// Here we wrap each node with #...# to avoid any edge cases when using strings.Contains
			return fmt.Sprintf("#%d%s%s#", node.Val, traversal(node.Left, "$"), traversal(node.Right, "%"))
		}
		return flag
	}
	return strings.Contains(traversal(root, ""), traversal(subRoot, ""))
}
