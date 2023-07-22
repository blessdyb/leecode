package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	sum := 0
	current := root
	for current != nil {
		if current.Left == nil {
			if current.Val >= low && current.Val <= high {
				sum += current.Val
			}
			current = current.Right
		} else {
			prev := current.Left
			for prev.Right != nil && prev.Right != current {
				prev = prev.Right
			}
			if prev.Right == nil {
				prev.Right = current
				current = current.Left
			}
			if prev.Right == current {
				prev.Right = nil
				if current.Val >= low && current.Val <= high {
					sum += current.Val
				}
				current = current.Right
			}
		}
	}
	return sum
}
