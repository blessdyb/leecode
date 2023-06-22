package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	var maxWithIndex = func(nums []int) (int, int) {
		max := -1
		index := -1
		for idx, item := range nums {
			if item > max {
				index = idx
				max = item
			}
		}
		return max, index
	}
	var recursive func(nums []int) *TreeNode
	recursive = func(nums []int) *TreeNode {
		var root *TreeNode
		if len(nums) > 0 {
			root = &TreeNode{}
			max, index := maxWithIndex(nums)
			root.Val = max
			left := nums[:index]
			right := nums[index+1:]
			root.Left = recursive(left)
			root.Right = recursive(right)
		}
		return root
	}
	return recursive(nums)
}
