package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	var ret *TreeNode
	if len(nums) > 0 {
		mid := (len(nums) - 1) / 2
		ret = &TreeNode{
			Val:   nums[mid],
			Left:  sortedArrayToBST(nums[:mid]),
			Right: sortedArrayToBST(nums[mid+1:]),
		}
	}
	return ret
}
