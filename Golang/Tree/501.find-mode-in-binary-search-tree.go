package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// If it's not a BST, we can traverse it and use a map to store the frequency
// Since it's a BST which results in a sorted list, so we just need to traverse the sorted array to get the max mode
func findMode(root *TreeNode) []int {
	ret := []int{}
	maxCount := 1
	previous := 1000000
	count := 1
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			if node.Val == previous {
				count++
			} else {
				count = 1
			}
			if maxCount == count {
				ret = append(ret, node.Val)
			} else if maxCount < count {
				maxCount = count
				ret = []int{node.Val}
			}
			previous = node.Val
			dfs(node.Right)
		}
	}
	dfs(root)
	return ret
}
