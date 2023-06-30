package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTreesRecursive(n int) []*TreeNode {
	var recursive func(start, end int) []*TreeNode
	recursive = func(start, end int) []*TreeNode {
		var ret []*TreeNode
		if end < start {
			ret = append(ret, nil)
			return ret
		}
		for i := start; i <= end; i++ {
			leftTrees := recursive(start, i-1)
			rightTrees := recursive(i+1, end)
			for _, itemL := range leftTrees {
				for _, itemR := range rightTrees {
					ret = append(ret, &TreeNode{
						Val:   i,
						Left:  itemL,
						Right: itemR,
					})
				}
			}
		}
		return ret
	}
	return recursive(1, n)
}
