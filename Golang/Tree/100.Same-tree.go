package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	var recursive func(p, q *TreeNode) bool
	recursive = func(p, q *TreeNode) bool {
		if (p == nil && q != nil) || (p != nil && q == nil) {
			return false
		} else if p == nil && q == nil {
			return true
		} else if p.Val != q.Val {
			return false
		}
		return recursive(p.Left, q.Left) && recursive(p.Right, q.Right)

	}
	return recursive(p, q)
}
