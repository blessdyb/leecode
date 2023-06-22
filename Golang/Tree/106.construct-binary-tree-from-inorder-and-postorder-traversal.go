package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(inorder []int, postorder []int) *TreeNode {
	var indexOf = func(slice []int, target int) int {
		for i, v := range slice {
			if v == target {
				return i
			}
		}
		return -1
	}
	var recursive func(i, p []int) *TreeNode
	recursive = func(i, p []int) *TreeNode {
		length := len(i)
		if length == 0 {
			return nil
		} else if length == 1 {
			return &TreeNode{
				Val:   i[0],
				Left:  nil,
				Right: nil,
			}
		}
		rootVal := p[length-1]
		spliterIndex := indexOf(i, rootVal)
		iLeft := i[:spliterIndex]
		iRight := i[spliterIndex+1:]
		pLeft := p[:len(iLeft)]
		pRight := p[len(iLeft) : len(p)-1]
		return &TreeNode{
			Val:   p[length-1],
			Left:  recursive(iLeft, pLeft),
			Right: recursive(iRight, pRight),
		}
	}
	return recursive(inorder, postorder)
}
