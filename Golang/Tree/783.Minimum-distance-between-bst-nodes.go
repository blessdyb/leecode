package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDiffInBST(root *TreeNode) int {
	ret := 100000
	var prev *TreeNode
	node := root
	stack := []*TreeNode{}
	for {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else if len(stack) > 0 {
			temp := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if prev != nil {
				diff := int(math.Abs(float64(prev.Val) - float64(temp.Val)))
				if diff < ret {
					ret = diff
				}
			}
			prev = temp
			node = temp.Right
		} else {
			break
		}
	}
	return ret
}
