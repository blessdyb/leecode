package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	} else if len(preorder) == 1 {
		return &TreeNode{Val: preorder[0]}
	}
	index := len(preorder)
	for i := 1; i < len(preorder); i++ {
		if preorder[i] > preorder[0] {
			index = i
			break
		}
	}
	return &TreeNode{
		Val:   preorder[0],
		Left:  bstFromPreorder(preorder[1:index]),
		Right: bstFromPreorder(preorder[index:]),
	}
}
