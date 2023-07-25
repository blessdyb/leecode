package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func allPossibleFBT(n int) []*TreeNode {
	cache := map[int][]*TreeNode{}
	var helper func(n int) []*TreeNode
	helper = func(n int) []*TreeNode {
		ret := []*TreeNode{}
		if n == 1 {
			ret = append(ret, &TreeNode{})
		} else if n%2 == 1 {
			if value, ok := cache[n]; ok {
				return value
			}
			for i := 1; i < n; i += 2 {
				left := helper(i)
				right := helper(n - 1 - i)
				for _, j := range left {
					for _, k := range right {
						root := &TreeNode{
							Left:  j,
							Right: k,
							Val:   0,
						}
						ret = append(ret, root)
					}
				}
			}
			cache[n] = ret
		}
		return ret
	}
	return helper(n)
}

func allPossibleFBTDP(n int) []*TreeNode {
	dp := make(map[int][]*TreeNode, n+1)
	dp[1] = []*TreeNode{&TreeNode{}}
	for i := 2; i <= n; i++ {
		if i%2 == 0 {
			dp[i] = []*TreeNode{}
		} else {
			nodes := []*TreeNode{}
			for j := 1; j < i; j += 2 {
				left := dp[j]
				right := dp[i-1-j]
				for _, k := range left {
					for _, l := range right {
						nodes = append(nodes, &TreeNode{
							Left:  k,
							Right: l,
						})
					}
				}
			}
			dp[i] = nodes
		}
	}
	return dp[n]
}
