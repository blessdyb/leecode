package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructFromPrePost(preorder []int, postorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	preorderWithoutRoot := preorder[1:]
	if len(preorderWithoutRoot) > 0 {
		postorderWithoutRoot := postorder[:len(preorderWithoutRoot)]
		rightChildVal := postorderWithoutRoot[len(postorderWithoutRoot)-1]
		preorderLeftNodes := []int{}
		for _, num := range preorderWithoutRoot {
			if num == rightChildVal {
				break
			}
			preorderLeftNodes = append(preorderLeftNodes, num)
		}
		preorderRightNodes := preorderWithoutRoot[len(preorderLeftNodes):]
		postorderLeftNodes := postorderWithoutRoot[:len(preorderLeftNodes)]
		postorderRightNodes := postorderWithoutRoot[len(postorderLeftNodes):]
		root.Left = constructFromPrePost(preorderLeftNodes, postorderLeftNodes)
		root.Right = constructFromPrePost(preorderRightNodes, postorderRightNodes)
	}
	return root
}
