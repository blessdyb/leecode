package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func pathSum(root *TreeNode, targetSum int) [][]int {
	ret := [][]int{}
	type TreeNodeWithPath struct {
		Node *TreeNode
		Path []int
	}
	var dfs func(node *TreeNodeWithPath, target int)
	dfs = func(node *TreeNodeWithPath, target int) {
		newPath := append(append([]int{}, node.Path...), node.Node.Val)
		if node.Node.Val == target && node.Node.Left == nil && node.Node.Right == nil {
			ret = append(ret, newPath)
		}
		if node.Node.Left != nil {
			dfs(&TreeNodeWithPath{
				Node: node.Node.Left,
				Path: newPath,
			}, target-node.Node.Val)
		}
		if node.Node.Right != nil {
			dfs(&TreeNodeWithPath{
				Node: node.Node.Right,
				Path: newPath,
			}, target-node.Node.Val)
		}
	}
	if root != nil {
		dfs(&TreeNodeWithPath{
			Node: root,
			Path: []int{},
		}, targetSum)
	}
	return ret
}
