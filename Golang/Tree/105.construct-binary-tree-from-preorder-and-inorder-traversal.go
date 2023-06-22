package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	var indexOf = func(haystack []int, target int) int {
		for index, item := range haystack {
			if item == target {
				return index
			}
		}
		return -1
	}
	var recursive func(p, i []int) *TreeNode
	recursive = func(p, i []int) *TreeNode {
		var root *TreeNode
		length := len(p)
		if length > 0 {
			root = &TreeNode{}
			rootValue := p[0]
			root.Val = rootValue
			rootValueIndex := indexOf(i, rootValue)
			iLeft := i[:rootValueIndex]
			iRight := i[rootValueIndex+1:]
			pLeft := p[1 : len(iLeft)+1]
			pRight := p[len(iLeft)+1:]
			root.Left = buildTree(pLeft, iLeft)
			root.Right = buildTree(pRight, iRight)
		}
		return root
	}
	return recursive(preorder, inorder)
}
