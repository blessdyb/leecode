package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func averageOfSubtree(root *TreeNode) int {
	ret := 0
	var helper func(node *TreeNode) (int, int)
	helper = func(node *TreeNode) (int, int) {
		if node == nil {
			return 0, 0
		}
		leftSum, leftNum := helper(node.Left)
		rightSum, rightNum := helper(node.Right)
		sum := node.Val + leftSum + rightSum
		num := leftNum + rightNum + 1
		if sum/num == node.Val {
			ret++
		}
		return sum, num
	}
	helper(root)
	return ret
}
