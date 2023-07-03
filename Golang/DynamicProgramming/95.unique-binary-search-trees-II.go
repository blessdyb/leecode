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

func generateTreesDP(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	dp := make([][]*TreeNode, n+1)
	dp[0] = []*TreeNode{nil}
	for i := 1; i <= n; i++ {
		dp[i] = []*TreeNode{}
		for j := 0; j < i; j++ {
			for _, treeL := range dp[j] {
				for _, treeR := range dp[i-j-1] {
					temp := &TreeNode{
						Val:   j + 1,
						Left:  treeL,
						Right: cloneTree(treeR, j+1),
					}
					dp[i] = append(dp[i], temp)
				}
			}
		}
	}
	return dp[n]
}

func cloneTree(root *TreeNode, offset int) *TreeNode {
	if root == nil {
		return root
	}
	temp := &TreeNode{}
	temp.Val = root.Val + offset
	temp.Left = cloneTree(root.Left, offset)
	temp.Right = cloneTree(root.Right, offset)
	return temp
}
